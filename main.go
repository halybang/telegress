package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/halybang/telegress/gitlab"
	"github.com/halybang/telegress/gogs"
	"github.com/halybang/telegress/jira"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var bots sync.Map

func RunBot(cfg *BotConfig, bot *tgbotapi.BotAPI) {

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {

	}

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}

}

func GitlabHandler(c *gin.Context) {
	body, ioerr := ioutil.ReadAll(c.Request.Body)
	if ioerr != nil {
		c.String(400, "Could not read request body")
		log.Println(ioerr)
		return
	}
	request, err := gitlab.Parse(string(body))
	if err != nil {
		c.String(400, "Could not parse request body")
		log.Println(err)
		return
	}
	log.Print(request)
}

func GogsHandler(c *gin.Context) {
	body, ioerr := ioutil.ReadAll(c.Request.Body)
	if ioerr != nil {
		c.String(400, "Could not read request body")
		log.Println(ioerr)
		return
	}
	request, err := gogs.Parse(string(body))
	if err != nil {
		c.String(400, "Could not parse request body")
		log.Println(err)
		return
	}
	log.Print(request)
}

func MakeHandleFunc(r *gin.Engine, config *BotConfig) gin.HandlerFunc {

	var bot *tgbotapi.BotAPI
	log.Printf("Init bot %v", config)
	result, ok := bots.Load(config.Token)
	if !ok || result == nil {
		//return nil
	} else {
		bot = result.(*tgbotapi.BotAPI)
	}
	if bot == nil {

	}
	cfg := *config

	wrappedHandler := func(r *gin.Engine, config *BotConfig) gin.HandlerFunc {
		return func(c *gin.Context) {
			var strNotify string
			var err error
			body, ioerr := ioutil.ReadAll(c.Request.Body)
			if ioerr != nil {
				c.String(400, "Could not read request body")
				log.Println(ioerr)
				return
			}

			//log.Printf("Handle request %v, data:%v", cfg.Uri, string(body))
			if cfg.Source == "gogs" {
				rq, err := gogs.Parse(string(body))
				if err != nil {
					log.Println(err)
				}
				strNotify = rq.String()
			} else if cfg.Source == "jira" {
				rq, err := jira.Parse(string(body))
				if err != nil {
					log.Println(err)
				} else {
					filename := strings.Replace(rq.WebhookEvent, ":", "_", -1) + ".txt"
					if _, err := os.Stat(filename); err != nil {
						f, err := os.Create(filename)
						if err != nil {
							log.Printf("Write jira request to file %s error:%v", filename, err)
						} else {
							log.Printf("Write jira request to file %s", filename)
							f.Write(body)
							f.Close()
						}
					}
					strNotify = rq.String()
				}
			}
			if err != nil {
				log.Printf("Error parse data: %v", err)
				c.JSON(200, "OK")
			}
			if strNotify == "" {
				c.JSON(200, "OK")
			}
			//log.Printf("Telebot output data: %s ", strNotify)
			// msg := tgbotapi.NewMessage(config.Channel, string(body))
			// msgRsp, errS := bot.Send(msg)
			// if errS != nil {
			// 	log.Printf("Telebot error: %v", errS)
			// } else {
			// 	log.Printf("Telebot response: %v", msgRsp)
			// }

			params := url.Values{}
			params.Add("chat_id", fmt.Sprintf("-%d", cfg.Channel))
			params.Add("text", strNotify)
			var url = fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", cfg.Token)
			//var str = fmt.Sprintf("chat_id=-%d&text=%s", cfg.Channel, string(strNotify))
			var str = params.Encode()
			log.Printf("Telebot send data: %s ", str)
			rsp, err := http.Post(url, "application/x-www-form-urlencoded", bytes.NewBuffer([]byte(str)))
			if err != nil {
				log.Printf("Telebot error: %v", err)
			} else {
				log.Printf("Telebot OK, %v", rsp)
			}
			c.JSON(200, "OK")
		}
	}(r, config)
	return wrappedHandler
}

func main() {

	flag.Int("port", 9092, "Listen Port")
	flag.String("address", "127.0.0.1", "Listen Address")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	viper.SetConfigType("json")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.BindPFlags(pflag.CommandLine)
	viper.ReadInConfig()

	config := &Config{
		Address: "127.0.0.1",
		Port:    9092,
	}
	config.Bots = []BotConfig{
		{
			Uri:      "gogs",
			Token:    "gogs",
			Format:   "gogs",
			Channel:  0,
			Commands: []string{},
		},
	}
	if err := viper.Unmarshal(&config); err != nil {
		log.Printf("couldn't read config: %s", err)
	}

	//gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	for _, botCfg := range config.Bots {
		var bot *tgbotapi.BotAPI
		result, ok := bots.Load(botCfg.Token)
		if !ok {
			botTmp, errBot := tgbotapi.NewBotAPI(botCfg.Token)
			if errBot != nil {
				log.Printf("Init bot token %s error %v", botCfg.Token, errBot)
				//continue
			} else {
				bot = botTmp
				bots.Store(botCfg.Token, bot)
			}
		} else {
			bot = result.(*tgbotapi.BotAPI)
			if bot == nil {
				//continue
			}
		}
		if bot != nil {
			log.Printf("Authorized on account %s", bot.Self.UserName)
		}
		// go RunBot(config, bot)
		r.POST(botCfg.Uri, MakeHandleFunc(r, &botCfg))
	}

	address := config.Address + ":" + strconv.FormatInt(config.Port, 10)
	r.Run(address)
}
