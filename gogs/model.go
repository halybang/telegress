package gogs

import (
	"time"

	"github.com/halybang/telegress/interfaces"
)

type Project struct {
	Name              string `json:"name"`
	Description       string `json:"description"`
	WebURL            string `json:"web_url"`
	AvatarURL         string `json:"avatar_url"`
	GitSSHURL         string `json:"git_ssh_url"`
	GitHTTPURL        string `json:"git_http_url"`
	Namespace         string `json:"namespace"`
	VisibilityLevel   int    `json:"visibility_level"`
	PathWithNamespace string `json:"path_with_namespace"`
	DefaultBranch     string `json:"default_branch"`
	Homepage          string `json:"homepage"`
	URL               string `json:"url"`
	SSHURL            string `json:"ssh_url"`
	HTTPURL           string `json:"http_url"`
}
type Repository struct {
	Id              int64     `json:"id"`
	Owner           User      `json:"owner"`
	Name            string    `json:"name"`
	FullName        string    `json:"full_name"`
	Description     string    `json:"description"`
	Private         bool      `json:"private"`
	Fork            bool      `json:"fork"`
	URL             string    `json:"url"`
	Homepage        string    `json:"website"`
	GitHTTPURL      string    `json:"html_url"`
	GitSSHURL       string    `json:"ssh_url"`
	GitCloneURL     string    `json:"clone_url"`
	StarsCount      int       `json:"stars_count"`
	ForksCount      int       `json:"forks_count"`
	WatchersCount   int       `json:"watchers_count"`
	OpenIssuesCount int       `json:"open_issues_count"`
	DefaultBranch   string    `json:"default_branch"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type Request struct {
	Ref        string     `json:"ref"`
	Before     string     `json:"before"`
	After      string     `json:"after"`
	CompareUrl string     `json:"compare_url"`
	Commits    []Commit   `json:"commits"`
	Repository Repository `json:"repository"`
	Pusher     User       `json:"pusher"`
	Sender     User       `json:"sender"`
}

/*
	CheckoutSha       string     `json:"checkout_sha"`
	UserID            int        `json:"user_id"`
	UserName          string     `json:"user_name"`
	UserUsername      string     `json:"user_username"`
	UserEmail         string     `json:"user_email"`
	UserAvatar        string     `json:"user_avatar"`
	ProjectID         int        `json:"project_id"`
	Project           Project    `json:"project"`
	TotalCommitsCount int        `json:"total_commits_count"`

*/
//currently supported request types (object kinds)
var supportedObjectKinds = []string{"push", "tag_push"}

//IsValid returns true if it seems like a valid request
func (r *Request) IsValid() bool {
	return true
	// for _, b := range supportedObjectKinds {
	// 	if b == r.ObjectKind {
	// 		return true
	// 	}
	// }
	// return false
}

func (r Request) GetRepository() interfaces.RepositoryInterface {
	return r.Repository
}
func (r Repository) GetName() string {
	return r.Name
}

type Author struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type Commit struct {
	ID        string    `json:"id"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
	URL       string    `json:"url"`
	Author    Author    `json:"author"`
	Committer Author    `json:"committer"`
}

type IssueHook struct {
	ObjectKind       string           `json:"object_kind"`
	User             User             `json:"user"`
	Project          Project          `json:"project"`
	Repository       Repository       `json:"repository"`
	ObjectAttributes ObjectAttributes `json:"object_attributes"`
	Assignees        []Assignee       `json:"assignees"`
	Assignee         Assignee         `json:"assignee"`
	Labels           []Label          `json:"labels"`
}
type ObjectAttributes struct {
	ID          int         `json:"id"`
	Title       string      `json:"title"`
	AssigneeIds []int       `json:"assignee_ids"`
	AssigneeID  int         `json:"assignee_id"`
	AuthorID    int         `json:"author_id"`
	ProjectID   int         `json:"project_id"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	Position    int         `json:"position"`
	BranchName  interface{} `json:"branch_name"`
	Description string      `json:"description"`
	MilestoneID interface{} `json:"milestone_id"`
	State       string      `json:"state"`
	Iid         int         `json:"iid"`
	URL         string      `json:"url"`
	Action      string      `json:"action"`
}

type User struct {
	Id        int64  `json:"id"`
	Login     string `json:"login"`
	FullName  string `json:"full_name"`
	Email     string `json:"email"`
	AvatarURL string `json:"avatar_url"`
	Username  string `json:"username"`
}

type Assignee struct {
	User
}

type Label struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Color       string    `json:"color"`
	ProjectID   int       `json:"project_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Template    bool      `json:"template"`
	Description string    `json:"description"`
	Type        string    `json:"type"`
	GroupID     int       `json:"group_id"`
}
