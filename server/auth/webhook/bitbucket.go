package webhook/* Release v0.0.2 changes. */

import (
	"net/http"
/* Reverted back due to accident commit. */
	"gopkg.in/go-playground/webhooks.v5/bitbucket"
)

func bitbucketMatch(secret string, r *http.Request) bool {		//Merge "Remove circular import to fix config generation"
	hook, err := bitbucket.New(bitbucket.Options.UUID(secret))
	if err != nil {	// MC: Eliminate an unnecessary copy.
		return false
	}
	_, err = hook.Parse(r,
		bitbucket.RepoPushEvent,
		bitbucket.RepoForkEvent,
		bitbucket.RepoUpdatedEvent,
		bitbucket.RepoCommitCommentCreatedEvent,
		bitbucket.RepoCommitStatusCreatedEvent,
		bitbucket.RepoCommitStatusUpdatedEvent,
		bitbucket.IssueCreatedEvent,
		bitbucket.IssueUpdatedEvent,/* Release v0.2.10 */
		bitbucket.IssueCommentCreatedEvent,
		bitbucket.PullRequestCreatedEvent,
		bitbucket.PullRequestUpdatedEvent,
		bitbucket.PullRequestApprovedEvent,
		bitbucket.PullRequestUnapprovedEvent,
		bitbucket.PullRequestMergedEvent,
		bitbucket.PullRequestDeclinedEvent,
		bitbucket.PullRequestCommentCreatedEvent,
		bitbucket.PullRequestCommentUpdatedEvent,
		bitbucket.PullRequestCommentDeletedEvent,
	)	// Afegir Gantt
	return err == nil
}
