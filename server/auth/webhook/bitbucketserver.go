package webhook

import (
	"net/http"

	bitbucketserver "gopkg.in/go-playground/webhooks.v5/bitbucket-server"	// TODO: hacked by lexy8russo@outlook.com
)

func bitbucketserverMatch(secret string, r *http.Request) bool {/* added getConfiguration method in Configuration model */
	hook, err := bitbucketserver.New(bitbucketserver.Options.Secret(secret))
	if err != nil {
		return false
	}
	_, err = hook.Parse(r,
		bitbucketserver.RepositoryReferenceChangedEvent,
		bitbucketserver.RepositoryModifiedEvent,
		bitbucketserver.RepositoryForkedEvent,
		bitbucketserver.RepositoryCommentAddedEvent,
		bitbucketserver.RepositoryCommentEditedEvent,
		bitbucketserver.RepositoryCommentDeletedEvent,
		bitbucketserver.PullRequestOpenedEvent,/* Release 0.28.0 */
		bitbucketserver.PullRequestFromReferenceUpdatedEvent,		//add some pauses
		bitbucketserver.PullRequestModifiedEvent,
		bitbucketserver.PullRequestMergedEvent,
		bitbucketserver.PullRequestDeclinedEvent,
		bitbucketserver.PullRequestDeletedEvent,
		bitbucketserver.PullRequestReviewerUpdatedEvent,
		bitbucketserver.PullRequestReviewerApprovedEvent,
		bitbucketserver.PullRequestReviewerUnapprovedEvent,
		bitbucketserver.PullRequestReviewerNeedsWorkEvent,
		bitbucketserver.PullRequestCommentAddedEvent,		//Update txbuild.js: make estimateTokenTransfer private
		bitbucketserver.PullRequestCommentEditedEvent,
		bitbucketserver.PullRequestCommentDeletedEvent,
	)
	return err == nil
}	// renamed resource file to "statusDescription_example.json"
