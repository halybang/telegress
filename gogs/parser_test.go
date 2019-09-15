package gogs

import (
	"reflect"
	"testing"
	"time"
)

var gogsPushRequest = `{
	"ref": "refs/heads/develop",
	"before": "28e1879d029cb852e4844d9c718537df08844e03",
	"after": "bffeb74224043ba2feb48d137756c8a9331c449a",
	"compare_url": "http://localhost:3000/unknwon/webhooks/compare/28e1879d029cb852e4844d9c718537df08844e03...bffeb74224043ba2feb48d137756c8a9331c449a",
	"commits": [{
		"id": "bffeb74224043ba2feb48d137756c8a9331c449a",
		"message": "!@#0^%\u003e\u003e\u003e\u003e\u003c\u003c\u003c\u003c\u003e\u003e\u003e\u003e\n",
		"url": "http://localhost:3000/unknwon/webhooks/commit/bffeb74224043ba2feb48d137756c8a9331c449a",
		"author": {
			"name": "Unknwon",
			"email": "u@gogs.io",
			"username": "unknwon"
		},
		"committer": {
			"name": "Unknwon",
			"email": "u@gogs.io",
			"username": "unknwon"
		},
		"timestamp": "2017-03-13T13:52:11-04:00"
	}],
	"repository": {
		"id": 140,
		"owner": {
			"id": 1,
			"login": "unknwon",
			"full_name": "Unknwon",
			"email": "u@gogs.io",
			"avatar_url": "https://secure.gravatar.com/avatar/d8b2871cdac01b57bbda23716cc03b96",
			"username": "unknwon"
		},
		"name": "webhooks",
		"full_name": "unknwon/webhooks",
		"description": "",
		"private": false,
		"fork": false,
		"html_url": "http://localhost:3000/unknwon/webhooks",
		"ssh_url": "ssh://unknwon@localhost:2222/unknwon/webhooks.git",
		"clone_url": "http://localhost:3000/unknwon/webhooks.git",
		"website": "",
		"stars_count": 0,
		"forks_count": 1,
		"watchers_count": 1,
		"open_issues_count": 7,
		"default_branch": "master",
		"created_at": "2017-02-26T04:29:06-05:00",
		"updated_at": "2017-03-13T13:51:58-04:00"
	},
	"pusher": {
		"id": 1,
		"login": "unknwon",
		"full_name": "Unknwon",
		"email": "u@gogs.io",
		"avatar_url": "https://secure.gravatar.com/avatar/d8b2871cdac01b57bbda23716cc03b96",
		"username": "unknwon"
	},
	"sender": {
		"id": 1,
		"login": "unknwon",
		"full_name": "Unknwon",
		"email": "u@gogs.io",
		"avatar_url": "https://secure.gravatar.com/avatar/d8b2871cdac01b57bbda23716cc03b96",
		"username": "unknwon"
	}
}`

func GogsSamplePushRequest() Request {
	str := "2017-02-26T04:29:06-05:00"
	t1, _ := time.Parse(time.RFC3339, str)
	t2, _ := time.Parse(time.RFC3339, "2017-03-13T13:51:58-04:00")
	t3, _ := time.Parse(time.RFC3339, "2017-03-13T13:52:11-04:00")
	return Request{
		Before:     "28e1879d029cb852e4844d9c718537df08844e03",
		After:      "bffeb74224043ba2feb48d137756c8a9331c449a",
		Ref:        "refs/heads/develop",
		CompareUrl: "http://localhost:3000/unknwon/webhooks/compare/28e1879d029cb852e4844d9c718537df08844e03...bffeb74224043ba2feb48d137756c8a9331c449a",

		Repository: Repository{
			Id: 140,
			Owner: User{
				Id:        1,
				Login:     "unknwon",
				FullName:  "Unknwon",
				Email:     "u@gogs.io",
				AvatarURL: "https://secure.gravatar.com/avatar/d8b2871cdac01b57bbda23716cc03b96",
				Username:  "unknwon",
			},
			Name:            "webhooks",
			FullName:        "unknwon/webhooks",
			Description:     "",
			Private:         false,
			Fork:            false,
			GitHTTPURL:      "http://localhost:3000/unknwon/webhooks",
			GitSSHURL:       "ssh://unknwon@localhost:2222/unknwon/webhooks.git",
			GitCloneURL:     "http://localhost:3000/unknwon/webhooks.git",
			Homepage:        "",
			StarsCount:      0,
			ForksCount:      1,
			WatchersCount:   1,
			OpenIssuesCount: 7,
			DefaultBranch:   "master",
			CreatedAt:       t1,
			UpdatedAt:       t2,
		},
		Commits: []Commit{
			Commit{
				ID:        "bffeb74224043ba2feb48d137756c8a9331c449a",
				Message:   "!@#0^%\u003e\u003e\u003e\u003e\u003c\u003c\u003c\u003c\u003e\u003e\u003e\u003e\n",
				Timestamp: t3,
				URL:       "http://localhost:3000/unknwon/webhooks/commit/bffeb74224043ba2feb48d137756c8a9331c449a",
				Author: Author{
					Name:     "Unknwon",
					Email:    "u@gogs.io",
					Username: "unknwon",
				},
				Committer: Author{
					Name:     "Unknwon",
					Email:    "u@gogs.io",
					Username: "unknwon",
				},
			},
		},
		Pusher: User{
			Id:        1,
			Login:     "unknwon",
			FullName:  "Unknwon",
			Email:     "u@gogs.io",
			AvatarURL: "https://secure.gravatar.com/avatar/d8b2871cdac01b57bbda23716cc03b96",
			Username:  "unknwon",
		},
		Sender: User{
			Id:        1,
			Login:     "unknwon",
			FullName:  "Unknwon",
			Email:     "u@gogs.io",
			AvatarURL: "https://secure.gravatar.com/avatar/d8b2871cdac01b57bbda23716cc03b96",
			Username:  "unknwon",
		},
	}
}
func TestParse(t *testing.T) {
	type args struct {
		jsonBody string
	}
	tests := []struct {
		name    string
		args    args
		wantR   Request
		wantErr bool
	}{
		{"push_sample", args{jsonBody: gogsPushRequest}, GogsSamplePushRequest(), false},
		{"push_syntax_error", args{jsonBody: "garbage"}, Request{}, true},
		{"push_empty", args{jsonBody: "{}"}, Request{}, true},
		//{"push_incomplete", args{jsonBody: `{"user_id": 4}`}, Request{UserID: 4}, true},
		//{"push_unsupported", args{jsonBody: `{"user_id": 4, "object_kind": "build"}`}, Request{UserID: 4, ObjectKind: "build"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotR, err := Parse(tt.args.jsonBody)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("Parse() = %v, want %v", gotR, tt.wantR)
			}

		})
	}
}
