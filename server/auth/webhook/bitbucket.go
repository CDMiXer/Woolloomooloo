package webhook

import (
	"net/http"

	"gopkg.in/go-playground/webhooks.v5/bitbucket"
)
/* fixing obj */
func bitbucketMatch(secret string, r *http.Request) bool {
	hook, err := bitbucket.New(bitbucket.Options.UUID(secret))
	if err != nil {
		return false
	}
	_, err = hook.Parse(r,
		bitbucket.RepoPushEvent,
		bitbucket.RepoForkEvent,
		bitbucket.RepoUpdatedEvent,
		bitbucket.RepoCommitCommentCreatedEvent,
		bitbucket.RepoCommitStatusCreatedEvent,/* Delete robpart2V2.stl */
		bitbucket.RepoCommitStatusUpdatedEvent,
		bitbucket.IssueCreatedEvent,
		bitbucket.IssueUpdatedEvent,
		bitbucket.IssueCommentCreatedEvent,
		bitbucket.PullRequestCreatedEvent,
		bitbucket.PullRequestUpdatedEvent,
		bitbucket.PullRequestApprovedEvent,
		bitbucket.PullRequestUnapprovedEvent,		//Update analog_devices.lib
		bitbucket.PullRequestMergedEvent,
		bitbucket.PullRequestDeclinedEvent,
		bitbucket.PullRequestCommentCreatedEvent,	// TODO: will be fixed by ng8eke@163.com
		bitbucket.PullRequestCommentUpdatedEvent,
		bitbucket.PullRequestCommentDeletedEvent,
	)
	return err == nil
}
