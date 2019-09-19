package jira

import (
	"time"

	"github.com/halybang/telegress/interfaces"
)

type Request struct {
	Timestamp    int64   `json:"timestamp"`
	WebhookEvent string  `json:"webhookEvent"`
	User         User    `json:"user"`
	Comment      Comment `json:"comment,omitempty"`
	Issue        Issue   `json:"issue,omitempty"`
	// After             string     `json:"after"`
	// Ref               string     `json:"ref"`
	// CheckoutSha       string     `json:"checkout_sha"`
	// UserID            int        `json:"user_id"`
	// UserName          string     `json:"user_name"`
	// UserUsername      string     `json:"user_username"`
	// UserEmail         string     `json:"user_email"`
	// UserAvatar        string     `json:"user_avatar"`
	// ProjectID         int        `json:"project_id"`
	// Project           Project    `json:"project"`
	// Repository        Repository `json:"repository"`
	// Commits           []Commit   `json:"commits"`
	// TotalCommitsCount int        `json:"total_commits_count"`
}

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
	//return r.Repository
	return nil
}
func (r Repository) GetName() string {
	return r.Name
}

type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Commit struct {
	ID        string   `json:"id"`
	Message   string   `json:"message"`
	Timestamp string   `json:"timestamp"`
	URL       string   `json:"url"`
	Author    Author   `json:"author"`
	Added     []string `json:"added"`
	Modified  []string `json:"modified"`
	Removed   []string `json:"removed"`
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
	Self        string `json:"self"`
	Name        string `json:"name,omitempty"`
	Email       string `json:"emailAddress,omitempty"`
	Key         string `json:"key,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	Active      bool   `json:"active,omitempty"`
	TimeZone    string `json:"timeZone,omitempty"`
	//AvatarURLs 	interface `json:"avatarUrls,omitempty"`
}

type Comment struct {
	Self         string `json:"self"`
	Id           string `json:"id"`
	Body         string `json:"body,omitempty"`
	Author       User   `json:"author,omitempty"`
	UpdateAuthor User   `json:"updateAuthor,omitempty"`
	Created      string `json:"created,omitempty"`
	Updated      string `json:"updated,omitempty"`
}

type Issue struct {
	Self   string     `json:"self"`
	Id     string     `json:"id"`
	Key    string     `json:"key,omitempty"`
	Fields IssueField `json:"fields,omitempty"`
}

type IssueField struct {
	IssueType   IssueType `json:"issuetype,omitempty"`
	Description string    `json:"description,omitempty"`
	Project     Project   `json:"project,omitempty"`
	WorkRatio   int64     `json:"workratio,omitempty"`
	Summary     string    `json:"summary,omitempty"`
	//Components []interface     `json:"components"`
	//FixVersions []interface     `json:"fixVersions"`

	//TimeSpent        interface    `json:"timespent"`
	//TimeEstimate        interface    `json:"timeestimate"`
	//TimeOriginalEstimate        interface    `json:"timeoriginalestimate"`
	//AggregateTimeOriginalEstimate        interface    `json:"aggregatetimeoriginalestimate"`

	//AggregateTimeSpent        interface    `json:"aggregatetimespent"`
	//AggregateTimeEstimate        interface    `json:"aggregatetimeestimate"`
	//AggregateProgress        interface    `json:"aggregateprogress"`

	//Resolution        interface    `json:"resolution"`
	//Timetracking        interface    `json:"timetracking"`
	//ResolutionDate        interface    `json:"resolutiondate"`

	LastViewed string `json:"lastViewed,omitempty"`
	Creator    User   `json:"creator,omitempty"`
	Reporter   User   `json:"reporter,omitempty"`
	// SubTask []interface `json:"subtasks,omitempty"`
	// Priority interface `json:"priority,omitempty"`
	// Labels []interface `json:"labels,omitempty"`
	// Environment interface `json:"environment,omitempty"`
	// Versions []interface `json:"versions,omitempty"`
	// DueDate interface `json:"duedate,omitempty"`
	// Progress interface `json:"progress,omitempty"`
	// Comment interface `json:"comment,omitempty"`
	// Issuelinks []interface `json:"issuelinks,omitempty"`
	// Votes interface `json:"votes,omitempty"`
	// WorkLog interface `json:"worklog,omitempty"`
	// Assignee interface `json:"assignee,omitempty"`
	Status  IssueStatus `json:"status,omitempty"`
	Created string      `json:"created,omitempty"`
	Updated string      `json:"updated,omitempty"`
}

type IssueType struct {
	Self        string      `json:"seft"`
	Id          string      `json:"id"`
	Description string      `json:"description"`
	IconUrl     string      `json:"iconUrl"`
	Name        string      `json:"name"`
	SubTask     interface{} `json:"subtask"`
}

type Project struct {
	Self           string `json:"self"`
	Id             string `json:"id"`
	Key            string `json:"key"`
	Name           string `json:"name"`
	ProjectTypeKey string `json:"projectTypeKey"`
}

type IssueStatus struct {
	Self        string `json:"seft"`
	Id          string `json:"id"`
	Description string `json:"description"`
	IconUrl     string `json:"iconUrl"`
	Name        string `json:"name"`
	//StatusCategory     interface `json:"statusCategory"`
}

type Repository struct {
	Name            string `json:"name"`
	URL             string `json:"url"`
	Description     string `json:"description"`
	Homepage        string `json:"homepage"`
	GitHTTPURL      string `json:"git_http_url"`
	GitSSHURL       string `json:"git_ssh_url"`
	VisibilityLevel int    `json:"visibility_level"`
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
