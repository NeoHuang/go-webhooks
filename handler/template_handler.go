package handler

import "github.com/google/go-github/github"

type TemplateHandler struct{}

func (handler *TemplateHandler) HandleCommitComment(event *github.CommitCommentEvent)       {}
func (handler *TemplateHandler) HandleCreate(event *github.CreateEvent)                     {}
func (handler *TemplateHandler) HandleDelete(event *github.DeleteEvent)                     {}
func (handler *TemplateHandler) HandleDeployment(event *github.DeploymentEvent)             {}
func (handler *TemplateHandler) HandleDeploymentStatus(event *github.DeploymentStatusEvent) {}
func (handler *TemplateHandler) HandleFork(event *github.ForkEvent)                         {}
func (handler *TemplateHandler) HandleGollum(event *github.GollumEvent)                     {}
func (handler *TemplateHandler) HandleIssueComment(event *github.IssueCommentEvent)         {}
func (handler *TemplateHandler) HandleIssues(event *github.IssueEvent)                      {}
func (handler *TemplateHandler) HandleMember(event *github.MemberEvent)                     {}
func (handler *TemplateHandler) HandleMembership(event *github.MembershipEvent)             {}
func (handler *TemplateHandler) HandlePageBuild(event *github.PageBuildEvent)               {}
func (handler *TemplateHandler) HandlePublic(event *github.PublicEvent)                     {}
func (handler *TemplateHandler) HandlePullRequest(event *github.PullRequestEvent)           {}
func (handler *TemplateHandler) HandlePullRequestReviewComment(event *github.PullRequestReviewCommentEvent) {
}
func (handler *TemplateHandler) HandlePush(event *github.PushEvent)             {}
func (handler *TemplateHandler) HandleRepository(event *github.RepositoryEvent) {}
func (handler *TemplateHandler) HandleRelease(event *github.ReleaseEvent)       {}
func (handler *TemplateHandler) HandleStatus(event *github.StatusEvent)         {}
func (handler *TemplateHandler) HandleTeamAdd(event *github.TeamAddEvent)       {}
func (handler *TemplateHandler) HandleWatch(event *github.WatchEvent)           {}
