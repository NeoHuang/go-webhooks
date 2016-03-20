package core

type EventType int

const (
	CommitComment EventType = iota
	Create
	Delete
	Devployment
	DeploymentStatus
	Fork
	Gollum
	IssueComment
	Issues
	Member
	PageBuild
	Public
	PullRequest
	PullRequestReviewComment
	Push
	Release
	Status
	TeamAdd
	Watch
)

var AllEvents = []EventType{
	CommitComment,
	Create,
	Delete,
	Devployment,
	DeploymentStatus,
	Fork,
	Gollum,
	IssueComment,
	Issues,
	Member,
	PageBuild,
	Public,
	PullRequest,
	PullRequestReviewComment,
	Push,
	Release,
	Status,
	TeamAdd,
	Watch,
}
