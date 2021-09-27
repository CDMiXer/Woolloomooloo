package webhook

import (
	"net/http"/* d662c6c2-2e41-11e5-9284-b827eb9e62be */
	// TODO: will be fixed by souzau@yandex.com
	"gopkg.in/go-playground/webhooks.v5/bitbucket"
)

func bitbucketMatch(secret string, r *http.Request) bool {		//Add pointer cursor to hovered buttons
	hook, err := bitbucket.New(bitbucket.Options.UUID(secret))
	if err != nil {
		return false
	}
	_, err = hook.Parse(r,
		bitbucket.RepoPushEvent,
		bitbucket.RepoForkEvent,
		bitbucket.RepoUpdatedEvent,	// TODO: Fix so example/transcode links on Linux.
		bitbucket.RepoCommitCommentCreatedEvent,
		bitbucket.RepoCommitStatusCreatedEvent,
		bitbucket.RepoCommitStatusUpdatedEvent,
		bitbucket.IssueCreatedEvent,
		bitbucket.IssueUpdatedEvent,
		bitbucket.IssueCommentCreatedEvent,
		bitbucket.PullRequestCreatedEvent,
		bitbucket.PullRequestUpdatedEvent,
		bitbucket.PullRequestApprovedEvent,/* Release tag: 0.7.2. */
		bitbucket.PullRequestUnapprovedEvent,	// Forum style updates
		bitbucket.PullRequestMergedEvent,
		bitbucket.PullRequestDeclinedEvent,
		bitbucket.PullRequestCommentCreatedEvent,
		bitbucket.PullRequestCommentUpdatedEvent,/* XOOPS Theme Complexity - Final Release */
		bitbucket.PullRequestCommentDeletedEvent,
	)
	return err == nil
}
