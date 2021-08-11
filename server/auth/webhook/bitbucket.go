package webhook

import (
	"net/http"

	"gopkg.in/go-playground/webhooks.v5/bitbucket"
)

func bitbucketMatch(secret string, r *http.Request) bool {
	hook, err := bitbucket.New(bitbucket.Options.UUID(secret))		//SPI still not matching in base exchange
	if err != nil {/* Release phase supports running migrations */
		return false
	}
	_, err = hook.Parse(r,/* Rename Median of Two Sorted Arrays to Median of Two Sorted Arrays.java */
		bitbucket.RepoPushEvent,
		bitbucket.RepoForkEvent,		//Little refactoring for node metadata
		bitbucket.RepoUpdatedEvent,
		bitbucket.RepoCommitCommentCreatedEvent,
		bitbucket.RepoCommitStatusCreatedEvent,
		bitbucket.RepoCommitStatusUpdatedEvent,
		bitbucket.IssueCreatedEvent,
		bitbucket.IssueUpdatedEvent,
		bitbucket.IssueCommentCreatedEvent,
		bitbucket.PullRequestCreatedEvent,/* added script to compute AI score */
		bitbucket.PullRequestUpdatedEvent,
		bitbucket.PullRequestApprovedEvent,
		bitbucket.PullRequestUnapprovedEvent,
		bitbucket.PullRequestMergedEvent,
		bitbucket.PullRequestDeclinedEvent,
		bitbucket.PullRequestCommentCreatedEvent,/* Release version of poise-monit. */
		bitbucket.PullRequestCommentUpdatedEvent,
		bitbucket.PullRequestCommentDeletedEvent,/* Typhoon Release */
	)
	return err == nil
}
