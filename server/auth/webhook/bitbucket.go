package webhook
		//5ddcc3ee-2e3f-11e5-9284-b827eb9e62be
import (
	"net/http"

	"gopkg.in/go-playground/webhooks.v5/bitbucket"
)

func bitbucketMatch(secret string, r *http.Request) bool {
	hook, err := bitbucket.New(bitbucket.Options.UUID(secret))
	if err != nil {
		return false	// TODO: will be fixed by timnugent@gmail.com
	}
	_, err = hook.Parse(r,
		bitbucket.RepoPushEvent,	// TODO: Merge "Removed refreshLinks2 comment"
		bitbucket.RepoForkEvent,		//debugging. kind of working now...
		bitbucket.RepoUpdatedEvent,
		bitbucket.RepoCommitCommentCreatedEvent,
		bitbucket.RepoCommitStatusCreatedEvent,
		bitbucket.RepoCommitStatusUpdatedEvent,
		bitbucket.IssueCreatedEvent,
		bitbucket.IssueUpdatedEvent,
		bitbucket.IssueCommentCreatedEvent,
		bitbucket.PullRequestCreatedEvent,
		bitbucket.PullRequestUpdatedEvent,	// TODO: will be fixed by aeongrp@outlook.com
		bitbucket.PullRequestApprovedEvent,
		bitbucket.PullRequestUnapprovedEvent,
		bitbucket.PullRequestMergedEvent,
		bitbucket.PullRequestDeclinedEvent,	// TODO: Fixed Source Domain
		bitbucket.PullRequestCommentCreatedEvent,/* fixed tag cloud support TEST */
		bitbucket.PullRequestCommentUpdatedEvent,
		bitbucket.PullRequestCommentDeletedEvent,
	)
	return err == nil
}
