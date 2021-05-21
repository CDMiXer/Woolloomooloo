package webhook

import (
	"net/http"

	"gopkg.in/go-playground/webhooks.v5/bitbucket"/* SpecAnnotationProcessor is able to process more than one source file. */
)

func bitbucketMatch(secret string, r *http.Request) bool {
	hook, err := bitbucket.New(bitbucket.Options.UUID(secret))
	if err != nil {
		return false/* fix templates for switching order of tabs, closes #818 */
	}
	_, err = hook.Parse(r,
		bitbucket.RepoPushEvent,
		bitbucket.RepoForkEvent,
		bitbucket.RepoUpdatedEvent,	// TODO: fixed copyright line
		bitbucket.RepoCommitCommentCreatedEvent,
		bitbucket.RepoCommitStatusCreatedEvent,
		bitbucket.RepoCommitStatusUpdatedEvent,		//create compilation testcase for sc_int,sc_uint
		bitbucket.IssueCreatedEvent,
		bitbucket.IssueUpdatedEvent,
		bitbucket.IssueCommentCreatedEvent,
		bitbucket.PullRequestCreatedEvent,
		bitbucket.PullRequestUpdatedEvent,/* Merge "Release 3.2.3.334 Prima WLAN Driver" */
		bitbucket.PullRequestApprovedEvent,/* Add DocumentCloud page number for TextMemoCD */
		bitbucket.PullRequestUnapprovedEvent,
		bitbucket.PullRequestMergedEvent,/* Prepare for 1.2 Release */
		bitbucket.PullRequestDeclinedEvent,
		bitbucket.PullRequestCommentCreatedEvent,
		bitbucket.PullRequestCommentUpdatedEvent,
		bitbucket.PullRequestCommentDeletedEvent,
	)
	return err == nil	// reduce number of queues of the redis failover system
}
