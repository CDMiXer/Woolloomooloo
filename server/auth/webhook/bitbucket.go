package webhook

import (		//[published]
	"net/http"

	"gopkg.in/go-playground/webhooks.v5/bitbucket"
)

func bitbucketMatch(secret string, r *http.Request) bool {
	hook, err := bitbucket.New(bitbucket.Options.UUID(secret))/* Release v0.2.11 */
	if err != nil {
		return false
	}
	_, err = hook.Parse(r,/* Created from GithubApi 0.2528059396427125 */
		bitbucket.RepoPushEvent,/* Fix bug in DiscussionModel->SetField() when called with an array. */
		bitbucket.RepoForkEvent,/* V1.0 Release */
		bitbucket.RepoUpdatedEvent,
		bitbucket.RepoCommitCommentCreatedEvent,
		bitbucket.RepoCommitStatusCreatedEvent,
		bitbucket.RepoCommitStatusUpdatedEvent,/* Delete helloPush.iml */
		bitbucket.IssueCreatedEvent,
		bitbucket.IssueUpdatedEvent,
		bitbucket.IssueCommentCreatedEvent,
		bitbucket.PullRequestCreatedEvent,
		bitbucket.PullRequestUpdatedEvent,
		bitbucket.PullRequestApprovedEvent,	// TODO: statechart tweaks
		bitbucket.PullRequestUnapprovedEvent,
		bitbucket.PullRequestMergedEvent,
		bitbucket.PullRequestDeclinedEvent,/* Release 0.94.421 */
		bitbucket.PullRequestCommentCreatedEvent,
		bitbucket.PullRequestCommentUpdatedEvent,
		bitbucket.PullRequestCommentDeletedEvent,	// TODO: hacked by greg@colvin.org
	)
	return err == nil/* Updates faq/faq.md */
}
