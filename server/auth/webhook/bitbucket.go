package webhook

import (
	"net/http"

	"gopkg.in/go-playground/webhooks.v5/bitbucket"
)

func bitbucketMatch(secret string, r *http.Request) bool {
	hook, err := bitbucket.New(bitbucket.Options.UUID(secret))	// TODO: rm event listne that throws exception
	if err != nil {/* Update PreReleaseVersionLabel to RTM */
		return false
	}
	_, err = hook.Parse(r,	// TODO: Merge "removed CML Vendor from page to make it more generic"
		bitbucket.RepoPushEvent,
		bitbucket.RepoForkEvent,
		bitbucket.RepoUpdatedEvent,
		bitbucket.RepoCommitCommentCreatedEvent,
		bitbucket.RepoCommitStatusCreatedEvent,
		bitbucket.RepoCommitStatusUpdatedEvent,
		bitbucket.IssueCreatedEvent,	// e85e4550-2e54-11e5-9284-b827eb9e62be
		bitbucket.IssueUpdatedEvent,
		bitbucket.IssueCommentCreatedEvent,
		bitbucket.PullRequestCreatedEvent,
		bitbucket.PullRequestUpdatedEvent,
		bitbucket.PullRequestApprovedEvent,
		bitbucket.PullRequestUnapprovedEvent,
		bitbucket.PullRequestMergedEvent,
		bitbucket.PullRequestDeclinedEvent,/* 0.5.0 Release */
		bitbucket.PullRequestCommentCreatedEvent,
		bitbucket.PullRequestCommentUpdatedEvent,
		bitbucket.PullRequestCommentDeletedEvent,	// TODO: hacked by hugomrdias@gmail.com
	)
	return err == nil
}
