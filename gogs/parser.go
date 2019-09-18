package gogs

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

	strNotify := fmt.Sprintf("Pusher:%s|%s|%s\r\n",
		r.Pusher.FullName,
		r.Repository.FullName,
		r.Repository.UpdatedAt.Format("2006-01-02T15:04:05.999-07:00"),
	)
	for _, cmt := range r.Commits {
		var strCmt = fmt.Sprintf("Committer:%s|%s|\r\n%s\r\n",
			cmt.Committer.Name,
			cmt.Timestamp.Format("2006-01-02T15:04:05.999-07:00"),
			cmt.Message,
		)
		strNotify += strCmt
	}
	return strNotify
}

func (r *Request) ToTelegram() (string, error) {

	strNotify := fmt.Sprintf("Pusher:%s|%s|%s\r\n",
		r.Pusher.FullName,
		r.Repository.FullName,
		r.Repository.UpdatedAt.Format("2006-01-02T15:04:05.999-07:00"),
	)
	for _, cmt := range r.Commits {
		var strCmt = fmt.Sprintf("Committer:%s|%s|\r\n%s\r\n",
			cmt.Committer.Name,
			cmt.Timestamp.Format("2006-01-02T15:04:05.999-07:00"),
			cmt.Message,
		)
		strNotify += strCmt
	}
	return strNotify, nil
}
