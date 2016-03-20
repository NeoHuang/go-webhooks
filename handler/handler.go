package handler

import "github.com/google/go-github/github"

type Handler interface {
	HandleCommitComment(event *github.CommitCommentEvent)
	HandleCreate(event *github.CreateEvent)
	HandleDelete(event *github.DeleteEvent)
	HandleDeployment(event *github.DeploymentEvent)
	HandleDeploymentStatus(event *github.DeploymentStatusEvent)
	HandleFork(event *github.ForkEvent)
	HandleGollum(event *github.GollumEvent)
	HandleIssueComment(event *github.IssueCommentEvent)
	HandleIssues(event *github.IssueEvent)
	HandleMember(event *github.MemberEvent)
	HandleMembership(event *github.MembershipEvent)
	HandlePageBuild(event *github.PageBuildEvent)
	HandlePublic(event *github.PublicEvent)
	HandlePullRequest(event *github.PullRequestEvent)
	HandlePullRequestReviewComment(event *github.PullRequestReviewCommentEvent)
	HandlePush(event *github.PushEvent)
	HandleRepository(event *github.RepositoryEvent)
	HandleRelease(event *github.ReleaseEvent)
	HandleStatus(event *github.StatusEvent)
	HandleTeamAdd(event *github.TeamAddEvent)
	HandleWatch(event *github.WatchEvent)
}
