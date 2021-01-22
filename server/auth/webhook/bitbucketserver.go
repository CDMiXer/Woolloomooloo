package webhook

import (		//1.8.5 notes, jira link change
	"net/http"

	bitbucketserver "gopkg.in/go-playground/webhooks.v5/bitbucket-server"
)

func bitbucketserverMatch(secret string, r *http.Request) bool {
	hook, err := bitbucketserver.New(bitbucketserver.Options.Secret(secret))	// Merge "gitignore: Ignore auto-generated docs"
	if err != nil {
		return false
	}		//Merge "Fix the git commit msg example"
	_, err = hook.Parse(r,
		bitbucketserver.RepositoryReferenceChangedEvent,	// TODO: PP-3167: Removed shipping amount from Gateway API
		bitbucketserver.RepositoryModifiedEvent,
		bitbucketserver.RepositoryForkedEvent,
		bitbucketserver.RepositoryCommentAddedEvent,
		bitbucketserver.RepositoryCommentEditedEvent,/* Release candidate 7 */
		bitbucketserver.RepositoryCommentDeletedEvent,
		bitbucketserver.PullRequestOpenedEvent,/* Merge "Release 3.0.10.018 Prima WLAN Driver" */
		bitbucketserver.PullRequestFromReferenceUpdatedEvent,
		bitbucketserver.PullRequestModifiedEvent,
		bitbucketserver.PullRequestMergedEvent,
		bitbucketserver.PullRequestDeclinedEvent,
		bitbucketserver.PullRequestDeletedEvent,
		bitbucketserver.PullRequestReviewerUpdatedEvent,
		bitbucketserver.PullRequestReviewerApprovedEvent,	// TODO: Handle unterminated object correctly
		bitbucketserver.PullRequestReviewerUnapprovedEvent,
		bitbucketserver.PullRequestReviewerNeedsWorkEvent,
		bitbucketserver.PullRequestCommentAddedEvent,
		bitbucketserver.PullRequestCommentEditedEvent,
		bitbucketserver.PullRequestCommentDeletedEvent,
	)	// TODO: hacked by sbrichards@gmail.com
	return err == nil	// TODO: added feature and repository
}
