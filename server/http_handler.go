package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/NeoHuang/go-webhooks/handler"
	"github.com/google/go-github/github"
)

type HttpHandler struct {
	eventHandler handler.Handler
}

type HandlingFunc func()

func NewHttpHandler(eventHandler handler.Handler) *HttpHandler {
	return &HttpHandler{
		eventHandler: eventHandler,
	}
}

func (handler *HttpHandler) ServeHTTP(writer http.ResponseWriter, httpRequest *http.Request) {
	fmt.Fprintf(writer, "OK")

	if handler.eventHandler == nil {
		log.Printf("error: no handler is set")
		return
	}

	if httpRequest.Header == nil {
		log.Printf("error: no header found")
		return
	}

	bodyBytes, err := ioutil.ReadAll(httpRequest.Body)
	if err != nil {
		log.Printf("error: failed to read request body")
		return
	}

	eventName := httpRequest.Header.Get("X-GitHub-Event")

	if fn := handler.getHandlingFunc(eventName, bodyBytes); fn != nil {
		go fn()
	}
}

func (handler *HttpHandler) getHandlingFunc(eventName string, bodyBytes []byte) HandlingFunc {
	switch eventName {
	case "commit_comment":
		event := &github.CommitCommentEvent{}
		err := json.Unmarshal(bodyBytes, event)
		if err != nil {
			logFailedEventMarshal(eventName)
		}
		return func() {
			handler.eventHandler.HandleCommitComment(event)
		}
	case "create":
		event := &github.CreateEvent{}
		err := json.Unmarshal(bodyBytes, event)
		if err != nil {
			logFailedEventMarshal(eventName)
		}
		return func() {
			handler.eventHandler.HandleCreate(event)
		}
	case "delete":
		event := &github.DeleteEvent{}
		err := json.Unmarshal(bodyBytes, event)
		if err != nil {
			logFailedEventMarshal(eventName)
		}
		return func() {
			handler.eventHandler.HandleDelete(event)
		}
	case "deployment":
		event := &github.DeploymentEvent{}
		err := json.Unmarshal(bodyBytes, event)
		if err != nil {
			logFailedEventMarshal(eventName)
		}
		return func() {
			handler.eventHandler.HandleDeployment(event)
		}
	case "deployment_status":
		event := &github.DeploymentStatusEvent{}
		err := json.Unmarshal(bodyBytes, event)
		if err != nil {
			logFailedEventMarshal(eventName)
		}
		return func() {
			handler.eventHandler.HandleDeploymentStatus(event)
		}
	case "fork":
		event := &github.ForkEvent{}
		err := json.Unmarshal(bodyBytes, event)
		if err != nil {
			logFailedEventMarshal(eventName)
		}
		return func() {
			handler.eventHandler.HandleFork(event)
		}
	case "gollum":
		event := &github.GollumEvent{}
		err := json.Unmarshal(bodyBytes, event)
		if err != nil {
			logFailedEventMarshal(eventName)
		}
		return func() {
			handler.eventHandler.HandleGollum(event)
		}
	case "issue_comment":
		event := &github.IssueCommentEvent{}
		err := json.Unmarshal(bodyBytes, event)
		if err != nil {
			logFailedEventMarshal(eventName)
		}
		return func() {
			handler.eventHandler.HandleIssueComment(event)
		}
	case "issues":
		event := &github.IssueEvent{}
		err := json.Unmarshal(bodyBytes, event)
		if err != nil {
			logFailedEventMarshal(eventName)
		}
		return func() {
			handler.eventHandler.HandleIssues(event)
		}
	case "member":
		event := &github.MemberEvent{}
		err := json.Unmarshal(bodyBytes, event)
		if err != nil {
			logFailedEventMarshal(eventName)
		}
		return func() {
			handler.eventHandler.HandleMember(event)
		}
	case "membership":
		event := &github.MembershipEvent{}
		err := json.Unmarshal(bodyBytes, event)
		if err != nil {
			logFailedEventMarshal(eventName)
		}
		return func() {
			handler.eventHandler.HandleMembership(event)
		}
	case "page_build":
		event := &github.PageBuildEvent{}
		err := json.Unmarshal(bodyBytes, event)
		if err != nil {
			logFailedEventMarshal(eventName)
		}
		return func() {
			handler.eventHandler.HandlePageBuild(event)
		}
	case "public":
		event := &github.PublicEvent{}
		err := json.Unmarshal(bodyBytes, event)
		if err != nil {
			logFailedEventMarshal(eventName)
		}
		return func() {
			handler.eventHandler.HandlePublic(event)
		}
	case "pull_request_review_comment":
		event := &github.PullRequestReviewCommentEvent{}
		err := json.Unmarshal(bodyBytes, event)
		if err != nil {
			logFailedEventMarshal(eventName)
		}
		return func() {
			handler.eventHandler.HandlePullRequestReviewComment(event)
		}
	case "pull_request":
		event := &github.PullRequestEvent{}
		err := json.Unmarshal(bodyBytes, event)
		if err != nil {
			logFailedEventMarshal(eventName)
		}
		return func() {
			handler.eventHandler.HandlePullRequest(event)
		}
	case "push":
		event := &github.PushEvent{}
		err := json.Unmarshal(bodyBytes, event)
		if err != nil {
			logFailedEventMarshal(eventName)
		}
		return func() {
			handler.eventHandler.HandlePush(event)
		}
	case "repository":
		event := &github.RepositoryEvent{}
		err := json.Unmarshal(bodyBytes, event)
		if err != nil {
			logFailedEventMarshal(eventName)
		}
		return func() {
			handler.eventHandler.HandleRepository(event)
		}
	case "release":
		event := &github.ReleaseEvent{}
		err := json.Unmarshal(bodyBytes, event)
		if err != nil {
			logFailedEventMarshal(eventName)
		}
		return func() {
			handler.eventHandler.HandleRelease(event)
		}
	case "status":
		event := &github.StatusEvent{}
		err := json.Unmarshal(bodyBytes, event)
		if err != nil {
			logFailedEventMarshal(eventName)
		}
		return func() {
			handler.eventHandler.HandleStatus(event)
		}
	case "team_add":
		event := &github.TeamAddEvent{}
		err := json.Unmarshal(bodyBytes, event)
		if err != nil {
			logFailedEventMarshal(eventName)
		}
		return func() {
			handler.eventHandler.HandleTeamAdd(event)
		}
	case "watch":
		event := &github.WatchEvent{}
		err := json.Unmarshal(bodyBytes, event)
		if err != nil {
			logFailedEventMarshal(eventName)
		}
		return func() {
			handler.eventHandler.HandleWatch(event)
		}
	}

	log.Printf("error: unknown event")
	return nil
}

func logFailedEventMarshal(eventName string) {
	log.Printf("error: failed to unmarshal event %s", eventName)
}
