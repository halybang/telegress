package jira

import (
	"reflect"
	"testing"
	"time"
)

var jiraPushRequest = `{"timestamp":1568604847751,"webhookEvent":"jira:issue_updated","issue_event_type_name":"issue_comment_edited","user":{"self":"http://jira.shlx.vn:8080/rest/api/2/user?username=tuanla","name":"tuanla","key":"tuanla","emailAddress":"tuanla@toanphuong.com.vn","avatarUrls":{"48x48":"http://jira.shlx.vn:8080/secure/useravatar?avatarId=10342","24x24":"http://jira.shlx.vn:8080/secure/useravatar?size=small&avatarId=10342","16x16":"http://jira.shlx.vn:8080/secure/useravatar?size=xsmall&avatarId=10342","32x32":"http://jira.shlx.vn:8080/secure/useravatar?size=medium&avatarId=10342"},"displayName":"Lưu Anh Tuấn","active":true,"timeZone":"Asia/Ho_Chi_Minh"},"issue":{"id":"10418","self":"http://jira.shlx.vn:8080/rest/api/2/issue/10418","key":"TECHSP-63","fields":{"issuetype":{"self":"http://jira.shlx.vn:8080/rest/api/2/issuetype/10001","id":"10001","description":"Created by Jira Software - do not edit or delete. Issue type for a user story.","iconUrl":"http://jira.shlx.vn:8080/images/icons/issuetypes/story.svg","name":"Story","subtask":false},"components":[],"timespent":null,"timeoriginalestimate":null,"description":"Chương trình MCS Nam Đàn bị lỗi kết nối với các xe: trên chương trình mcs báo các xe có kết nối và hiện IP, ping vào địa chỉ IP đó OK, nhưng từ chương trình MCS không chụp ảnh ảnh và cho xe chạy được. Khi bị thường phải chờ một lúc rồi khởi động lại MCS thì có thể lại bình thường.","project":{"self":"http://jira.shlx.vn:8080/rest/api/2/project/10001","id":"10001","key":"TECHSP","name":"Techsupport","projectTypeKey":"software","avatarUrls":{"48x48":"http://jira.shlx.vn:8080/secure/projectavatar?avatarId=10324","24x24":"http://jira.shlx.vn:8080/secure/projectavatar?size=small&avatarId=10324","16x16":"http://jira.shlx.vn:8080/secure/projectavatar?size=xsmall&avatarId=10324","32x32":"http://jira.shlx.vn:8080/secure/projectavatar?size=medium&avatarId=10324"}},"fixVersions":[],"aggregatetimespent":null,"resolution":null,"timetracking":{},"customfield_10105":null,"customfield_10106":null,"customfield_10107":{"self":"http://jira.shlx.vn:8080/rest/api/2/customFieldOption/10033","value":"Nghệ An-Nam Đàn-A1","id":"10033"},"attachment":[],"aggregatetimeestimate":null,"resolutiondate":null,"workratio":-1,"summary":"Lỗi phần phần mềm MCS A1 nam đàn","lastViewed":"2019-09-16T10:34:07.711+0700","watches":{"self":"http://jira.shlx.vn:8080/rest/api/2/issue/TECHSP-63/watchers","watchCount":2,"isWatching":true},"creator":{"self":"http://jira.shlx.vn:8080/rest/api/2/user?username=tuanla","name":"tuanla","key":"tuanla","emailAddress":"tuanla@toanphuong.com.vn","avatarUrls":{"48x48":"http://jira.shlx.vn:8080/secure/useravatar?avatarId=10342","24x24":"http://jira.shlx.vn:8080/secure/useravatar?size=small&avatarId=10342","16x16":"http://jira.shlx.vn:8080/secure/useravatar?size=xsmall&avatarId=10342","32x32":"http://jira.shlx.vn:8080/secure/useravatar?size=medium&avatarId=10342"},"displayName":"Lưu Anh Tuấn","active":true,"timeZone":"Asia/Ho_Chi_Minh"},"subtasks":[],"created":"2019-09-16T08:57:21.314+0700","reporter":{"self":"http://jira.shlx.vn:8080/rest/api/2/user?username=tuanla","name":"tuanla","key":"tuanla","emailAddress":"tuanla@toanphuong.com.vn","avatarUrls":{"48x48":"http://jira.shlx.vn:8080/secure/useravatar?avatarId=10342","24x24":"http://jira.shlx.vn:8080/secure/useravatar?size=small&avatarId=10342","16x16":"http://jira.shlx.vn:8080/secure/useravatar?size=xsmall&avatarId=10342","32x32":"http://jira.shlx.vn:8080/secure/useravatar?size=medium&avatarId=10342"},"displayName":"Lưu Anh Tuấn","active":true,"timeZone":"Asia/Ho_Chi_Minh"},"customfield_10000":"{summaryBean=com.atlassian.jira.plugin.devstatus.rest.SummaryBean@724a043b[summary={pullrequest=com.atlassian.jira.plugin.devstatus.rest.SummaryItemBean@7e8fca8c[overall=PullRequestOverallBean{stateCount=0, state='OPEN', details=PullRequestOverallDetails{openCount=0, mergedCount=0, declinedCount=0}},byInstanceType={}], build=com.atlassian.jira.plugin.devstatus.rest.SummaryItemBean@54a48937[overall=com.atlassian.jira.plugin.devstatus.summary.beans.BuildOverallBean@547ff285[failedBuildCount=0,successfulBuildCount=0,unknownBuildCount=0,count=0,lastUpdated=<null>,lastUpdatedTimestamp=<null>],byInstanceType={}], review=com.atlassian.jira.plugin.devstatus.rest.SummaryItemBean@100e1fa[overall=com.atlassian.jira.plugin.devstatus.summary.beans.ReviewsOverallBean@1ef2440f[stateCount=0,state=<null>,dueDate=<null>,overDue=false,count=0,lastUpdated=<null>,lastUpdatedTimestamp=<null>],byInstanceType={}], deployment-environment=com.atlassian.jira.plugin.devstatus.rest.SummaryItemBean@6e6f8bd7[overall=com.atlassian.jira.plugin.devstatus.summary.beans.DeploymentOverallBean@452fa745[topEnvironments=[],showProjects=false,successfulCount=0,count=0,lastUpdated=<null>,lastUpdatedTimestamp=<null>],byInstanceType={}], repository=com.atlassian.jira.plugin.devstatus.rest.SummaryItemBean@4326912b[overall=com.atlassian.jira.plugin.devstatus.summary.beans.CommitOverallBean@4ff8777c[count=0,lastUpdated=<null>,lastUpdatedTimestamp=<null>],byInstanceType={}], branch=com.atlassian.jira.plugin.devstatus.rest.SummaryItemBean@5d5b7c67[overall=com.atlassian.jira.plugin.devstatus.summary.beans.BranchOverallBean@6576fd06[count=0,lastUpdated=<null>,lastUpdatedTimestamp=<null>],byInstanceType={}]},errors=[],configErrors=[]], devSummaryJson={\"cachedValue\":{\"errors\":[],\"configErrors\":[],\"summary\":{\"pullrequest\":{\"overall\":{\"count\":0,\"lastUpdated\":null,\"stateCount\":0,\"state\":\"OPEN\",\"details\":{\"openCount\":0,\"mergedCount\":0,\"declinedCount\":0,\"total\":0},\"open\":true},\"byInstanceType\":{}},\"build\":{\"overall\":{\"count\":0,\"lastUpdated\":null,\"failedBuildCount\":0,\"successfulBuildCount\":0,\"unknownBuildCount\":0},\"byInstanceType\":{}},\"review\":{\"overall\":{\"count\":0,\"lastUpdated\":null,\"stateCount\":0,\"state\":null,\"dueDate\":null,\"overDue\":false,\"completed\":false},\"byInstanceType\":{}},\"deployment-environment\":{\"overall\":{\"count\":0,\"lastUpdated\":null,\"topEnvironments\":[],\"showProjects\":false,\"successfulCount\":0},\"byInstanceType\":{}},\"repository\":{\"overall\":{\"count\":0,\"lastUpdated\":null},\"byInstanceType\":{}},\"branch\":{\"overall\":{\"count\":0,\"lastUpdated\":null},\"byInstanceType\":{}}}},\"isStale\":false}}","aggregateprogress":{"progress":0,"total":0},"priority":{"self":"http://jira.shlx.vn:8080/rest/api/2/priority/3","iconUrl":"http://jira.shlx.vn:8080/images/icons/priorities/medium.svg","name":"Medium","id":"3"},"customfield_10100":"0|i000f3:","customfield_10101":null,"labels":[],"environment":null,"timeestimate":null,"aggregatetimeoriginalestimate":null,"versions":[],"duedate":null,"progress":{"progress":0,"total":0},"comment":{"comments":[{"self":"http://jira.shlx.vn:8080/rest/api/2/issue/10418/comment/10453","id":"10453","author":{"self":"http://jira.shlx.vn:8080/rest/api/2/user?username=quanpl","name":"quanpl","key":"quanpl","emailAddress":"quanpl@toanphuong.com.vn","avatarUrls":{"48x48":"https://www.gravatar.com/avatar/71270706ae787e16a899340b205fd27c?d=mm&s=48","24x24":"https://www.gravatar.com/avatar/71270706ae787e16a899340b205fd27c?d=mm&s=24","16x16":"https://www.gravatar.com/avatar/71270706ae787e16a899340b205fd27c?d=mm&s=16","32x32":"https://www.gravatar.com/avatar/71270706ae787e16a899340b205fd27c?d=mm&s=32"},"displayName":"Phan Lương Quân","active":true,"timeZone":"Asia/Ho_Chi_Minh"},"body":"Anh [~admin] kiểm tra giúp nhé, họ báo là trước mỗi lần thi đều phải nhờ anh làm gì đó mới được.","updateAuthor":{"self":"http://jira.shlx.vn:8080/rest/api/2/user?username=quanpl","name":"quanpl","key":"quanpl","emailAddress":"quanpl@toanphuong.com.vn","avatarUrls":{"48x48":"https://www.gravatar.com/avatar/71270706ae787e16a899340b205fd27c?d=mm&s=48","24x24":"https://www.gravatar.com/avatar/71270706ae787e16a899340b205fd27c?d=mm&s=24","16x16":"https://www.gravatar.com/avatar/71270706ae787e16a899340b205fd27c?d=mm&s=16","32x32":"https://www.gravatar.com/avatar/71270706ae787e16a899340b205fd27c?d=mm&s=32"},"displayName":"Phan Lương Quân","active":true,"timeZone":"Asia/Ho_Chi_Minh"},"created":"2019-09-16T09:05:23.234+0700","updated":"2019-09-16T09:05:23.234+0700"},{"self":"http://jira.shlx.vn:8080/rest/api/2/issue/10418/comment/10454","id":"10454","author":{"self":"http://jira.shlx.vn:8080/rest/api/2/user?username=tuanla","name":"tuanla","key":"tuanla","emailAddress":"tuanla@toanphuong.com.vn","avatarUrls":{"48x48":"http://jira.shlx.vn:8080/secure/useravatar?avatarId=10342","24x24":"http://jira.shlx.vn:8080/secure/useravatar?size=small&avatarId=10342","16x16":"http://jira.shlx.vn:8080/secure/useravatar?size=xsmall&avatarId=10342","32x32":"http://jira.shlx.vn:8080/secure/useravatar?size=medium&avatarId=10342"},"displayName":"Lưu Anh Tuấn","active":true,"timeZone":"Asia/Ho_Chi_Minh"},"body":"Anh Hà dự đoán nguyên nhân là gì? (ở Nam Đàn là dùng bản MCS mới nhất anh đưa)","updateAuthor":{"self":"http://jira.shlx.vn:8080/rest/api/2/user?username=tuanla","name":"tuanla","key":"tuanla","emailAddress":"tuanla@toanphuong.com.vn","avatarUrls":{"48x48":"http://jira.shlx.vn:8080/secure/useravatar?avatarId=10342","24x24":"http://jira.shlx.vn:8080/secure/useravatar?size=small&avatarId=10342","16x16":"http://jira.shlx.vn:8080/secure/useravatar?size=xsmall&avatarId=10342","32x32":"http://jira.shlx.vn:8080/secure/useravatar?size=medium&avatarId=10342"},"displayName":"Lưu Anh Tuấn","active":true,"timeZone":"Asia/Ho_Chi_Minh"},"created":"2019-09-16T10:21:10.277+0700","updated":"2019-09-16T10:21:10.277+0700"},{"self":"http://jira.shlx.vn:8080/rest/api/2/issue/10418/comment/10455","id":"10455","author":{"self":"http://jira.shlx.vn:8080/rest/api/2/user?username=tuanla","name":"tuanla","key":"tuanla","emailAddress":"tuanla@toanphuong.com.vn","avatarUrls":{"48x48":"http://jira.shlx.vn:8080/secure/useravatar?avatarId=10342","24x24":"http://jira.shlx.vn:8080/secure/useravatar?size=small&avatarId=10342","16x16":"http://jira.shlx.vn:8080/secure/useravatar?size=xsmall&avatarId=10342","32x32":"http://jira.shlx.vn:8080/secure/useravatar?size=medium&avatarId=10342"},"displayName":"Lưu Anh Tuấn","active":true,"timeZone":"Asia/Ho_Chi_Minh"},"body":"Các sân khác : A1 - Thủy Lợi, Tư Thục - Bình Phước lại không bị hiện tượng như vậy.","updateAuthor":{"self":"http://jira.shlx.vn:8080/rest/api/2/user?username=tuanla","name":"tuanla","key":"tuanla","emailAddress":"tuanla@toanphuong.com.vn","avatarUrls":{"48x48":"http://jira.shlx.vn:8080/secure/useravatar?avatarId=10342","24x24":"http://jira.shlx.vn:8080/secure/useravatar?size=small&avatarId=10342","16x16":"http://jira.shlx.vn:8080/secure/useravatar?size=xsmall&avatarId=10342","32x32":"http://jira.shlx.vn:8080/secure/useravatar?size=medium&avatarId=10342"},"displayName":"Lưu Anh Tuấn","active":true,"timeZone":"Asia/Ho_Chi_Minh"},"created":"2019-09-16T10:25:42.187+0700","updated":"2019-09-16T10:34:07.715+0700"}],"maxResults":3,"total":3,"startAt":0},"issuelinks":[],"votes":{"self":"http://jira.shlx.vn:8080/rest/api/2/issue/TECHSP-63/votes","votes":0,"hasVoted":false},"worklog":{"startAt":0,"maxResults":20,"total":0,"worklogs":[]},"assignee":{"self":"http://jira.shlx.vn:8080/rest/api/2/user?username=admin","name":"admin","key":"admin","emailAddress":"dinhha@toanphuong.com.vn","avatarUrls":{"48x48":"https://www.gravatar.com/avatar/ce99592c7cd7083a628b61def4590263?d=mm&s=48","24x24":"https://www.gravatar.com/avatar/ce99592c7cd7083a628b61def4590263?d=mm&s=24","16x16":"https://www.gravatar.com/avatar/ce99592c7cd7083a628b61def4590263?d=mm&s=16","32x32":"https://www.gravatar.com/avatar/ce99592c7cd7083a628b61def4590263?d=mm&s=32"},"displayName":"Dinh Ha","active":true,"timeZone":"Asia/Ho_Chi_Minh"},"updated":"2019-09-16T10:29:50.823+0700","status":{"self":"http://jira.shlx.vn:8080/rest/api/2/status/10002","description":"","iconUrl":"http://jira.shlx.vn:8080/","name":"Backlog","id":"10002","statusCategory":{"self":"http://jira.shlx.vn:8080/rest/api/2/statuscategory/2","id":2,"key":"new","colorName":"blue-gray","name":"To Do"}}}},"comment":{"self":"http://jira.shlx.vn:8080/rest/api/2/issue/10418/comment/10455","id":"10455","author":{"self":"http://jira.shlx.vn:8080/rest/api/2/user?username=tuanla","name":"tuanla","key":"tuanla","emailAddress":"tuanla@toanphuong.com.vn","avatarUrls":{"48x48":"http://jira.shlx.vn:8080/secure/useravatar?avatarId=10342","24x24":"http://jira.shlx.vn:8080/secure/useravatar?size=small&avatarId=10342","16x16":"http://jira.shlx.vn:8080/secure/useravatar?size=xsmall&avatarId=10342","32x32":"http://jira.shlx.vn:8080/secure/useravatar?size=medium&avatarId=10342"},"displayName":"Lưu Anh Tuấn","active":true,"timeZone":"Asia/Ho_Chi_Minh"},"body":"Các sân khác : A1 - Thủy Lợi, Tư Thục - Bình Phước lại không bị hiện tượng như vậy.","updateAuthor":{"self":"http://jira.shlx.vn:8080/rest/api/2/user?username=tuanla","name":"tuanla","key":"tuanla","emailAddress":"tuanla@toanphuong.com.vn","avatarUrls":{"48x48":"http://jira.shlx.vn:8080/secure/useravatar?avatarId=10342","24x24":"http://jira.shlx.vn:8080/secure/useravatar?size=small&avatarId=10342","16x16":"http://jira.shlx.vn:8080/secure/useravatar?size=xsmall&avatarId=10342","32x32":"http://jira.shlx.vn:8080/secure/useravatar?size=medium&avatarId=10342"},"displayName":"Lưu Anh Tuấn","active":true,"timeZone":"Asia/Ho_Chi_Minh"},"created":"2019-09-16T10:25:42.187+0700","updated":"2019-09-16T10:34:07.715+0700"}}`

func GitlabSamplePushRequest() Request {
	str := "2011-12-12T14:27:31+02:00"
	t1, _ := time.Parse(time.RFC3339, str)
	t2, _ := time.Parse(time.RFC3339, "2012-01-03T23:36:29+02:00")

	return Request{
		ObjectKind:   "push",
		Before:       "95790bf891e76fee5e1747ab589903a6a1f80f22",
		After:        "da1560886d4f094c3e6c9ef40349f7d38b5d27d7",
		Ref:          "refs/heads/master",
		CheckoutSha:  "da1560886d4f094c3e6c9ef40349f7d38b5d27d7",
		UserID:       4,
		UserName:     "John Smith",
		UserUsername: "jsmith",
		UserEmail:    "john@example.com",
		UserAvatar:   "https://s.gravatar.com/avatar/d4c74594d841139328695756648b6bd6?s=8://s.gravatar.com/avatar/d4c74594d841139328695756648b6bd6?s=80",
		ProjectID:    15,
		Project: Project{
			Name:              "Diaspora",
			Description:       "",
			WebURL:            "http://example.com/mike/diaspora",
			GitSSHURL:         "git@example.com:mike/diaspora.git",
			GitHTTPURL:        "http://example.com/mike/diaspora.git",
			Namespace:         "Mike",
			VisibilityLevel:   0,
			PathWithNamespace: "mike/diaspora",
			DefaultBranch:     "master",
			Homepage:          "http://example.com/mike/diaspora",
			URL:               "git@example.com:mike/diaspora.git",
			SSHURL:            "git@example.com:mike/diaspora.git",
			HTTPURL:           "http://example.com/mike/diaspora.git",
		},
		Repository: Repository{
			Name:            "Diaspora",
			URL:             "git@example.com:mike/diaspora.git",
			Description:     "",
			Homepage:        "http://example.com/mike/diaspora",
			GitHTTPURL:      "http://example.com/mike/diaspora.git",
			GitSSHURL:       "git@example.com:mike/diaspora.git",
			VisibilityLevel: 0,
		},
		Commits: []Commit{
			Commit{
				ID:        "b6568db1bc1dcd7f8b4d5a946b0b91f9dacd7327",
				Message:   "Update Catalan translation to e38cb41.",
				Timestamp: t1,
				URL:       "http://example.com/mike/diaspora/commit/b6568db1bc1dcd7f8b4d5a946b0b91f9dacd7327",
				Author: Author{
					Name:  "Jordi Mallach",
					Email: "jordi@softcatala.org",
				},
				Added:    []string{"CHANGELOG"},
				Modified: []string{"app/controller/application.rb"},
				Removed:  []string{},
			},
			Commit{
				ID:        "da1560886d4f094c3e6c9ef40349f7d38b5d27d7",
				Message:   "fixed readme",
				Timestamp: t2,
				URL:       "http://example.com/mike/diaspora/commit/da1560886d4f094c3e6c9ef40349f7d38b5d27d7",
				Author: Author{
					Name:  "GitLab dev user",
					Email: "gitlabdev@dv6700.(none)",
				},
				Added:    []string{"CHANGELOG"},
				Modified: []string{"app/controller/application.rb"},
				Removed:  []string{},
			},
		},
		TotalCommitsCount: 4,
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
		{"push_sample", args{jsonBody: gitlabPushRequest}, GitlabSamplePushRequest(), false},
		{"push_syntax_error", args{jsonBody: "garbage"}, Request{}, true},
		{"push_empty", args{jsonBody: "{}"}, Request{}, true},
		{"push_incomplete", args{jsonBody: `{"user_id": 4}`}, Request{UserID: 4}, true},
		{"push_unsupported", args{jsonBody: `{"user_id": 4, "object_kind": "build"}`}, Request{UserID: 4, ObjectKind: "build"}, true},
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
