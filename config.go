package main

//"encoding/json"
//	"errors"
//	"log"
//	"os"
//	"os/signal"
//	"syscall"

type Config struct {
	Address string      `json:"Address"`
	Port    int64       `json:"Port"`
	Bots    []BotConfig `json:"Bots"`
}

type BotConfig struct {
	Uri      string   `json:"uri"`
	Token    string   `json:"token"`
	Format   string   `json:"fmt"`
	Source   string   `json:"source"`
	Channel  int64    `json:"channel"`
	Commands []string `json:"commands"`
}
