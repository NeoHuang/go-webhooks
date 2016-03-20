package handler

import (
	"log"

	"github.com/google/go-github/github"
)

type LogHandler struct {
	TemplateHandler
}

func NewLogHandler() *LogHandler {
	return &LogHandler{}
}

func (handler *LogHandler) HandleCommitComment(event *github.CommitCommentEvent) {}

func (handler *LogHandler) HandleCreate(event *github.CreateEvent) {}

func (handler *LogHandler) HandleDelete(event *github.DeleteEvent) {}

func (handler *LogHandler) HandleDeployment(event *github.DeploymentEvent) {}

func (handler *LogHandler) HandleDeploymentStatus(event *github.DeploymentStatusEvent) {}

func (handler *LogHandler) HandleFork(event *github.ForkEvent) {}

func (handler *LogHandler) HandleGollum(event *github.GollumEvent) {}

func (handler *LogHandler) HandleIssueComment(event *github.IssueCommentEvent) {
	log.Printf("issue commented: %q by %q", *event.Comment.Body, *event.Comment.User.Login)
}

func (handler *LogHandler) HandleIssues(event *github.IssueEvent) {}

func (handler *LogHandler) HandleMember(event *github.MemberEvent) {}

func (handler *LogHandler) HandleMembership(event *github.MembershipEvent) {}

func (handler *LogHandler) HandlePageBuild(event *github.PageBuildEvent) {}

func (handler *LogHandler) HandlePublic(event *github.PublicEvent) {}

func (handler *LogHandler) HandlePullRequest(event *github.PullRequestEvent) {}

func (handler *LogHandler) HandlePullRequestReviewComment(event *github.PullRequestReviewCommentEvent) {
}

func (handler *LogHandler) HandlePush(event *github.PushEvent) {}

func (handler *LogHandler) HandleRepository(event *github.RepositoryEvent) {}

func (handler *LogHandler) HandleRelease(event *github.ReleaseEvent) {}

func (handler *LogHandler) HandleStatus(event *github.StatusEvent) {}

func (handler *LogHandler) HandleTeamAdd(event *github.TeamAddEvent) {}

func (handler *LogHandler) HandleWatch(event *github.WatchEvent) {}
