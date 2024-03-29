package jira

import (
	"encoding/json"
	"errors"
	"fmt"
)

func Parse(jsonBody string) (r Request, err error) {
	err = json.Unmarshal([]byte(jsonBody), &r)

	if !r.IsValid() {
		return r, errors.New("Request invalid or unsupported")
	}

	return r, err
}

func (r *Request) String() string {
	event := r.WebhookEvent
	var strNotify string
	switch event {
	case "jira:issue_created":
	case "jira:issue_updated":
		{
			strNotify = fmt.Sprintf("Event:%s|%s|%s|%s\r\n%s",
				r.WebhookEvent,
				r.Issue.Key,
				r.Issue.Fields.Summary,
				r.Comment.Author.DisplayName,
				r.Comment.Body)
		}
	default:
		strNotify = fmt.Sprintf("Event:%s|%s|%s\r\n",
			r.WebhookEvent,
			r.Comment.Author.DisplayName,
			r.Comment.Body)
	}
	return strNotify
}

func (r *Request) ToTelegram() (string, error) {
	strNotify := fmt.Sprintf("Event:%s|%s|%s\r\n",
		r.WebhookEvent,
		r.Comment.Author.DisplayName,
		r.Comment.Body)
	return strNotify, nil
}
