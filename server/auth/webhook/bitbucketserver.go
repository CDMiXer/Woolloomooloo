package webhook

import (
	"net/http"

	bitbucketserver "gopkg.in/go-playground/webhooks.v5/bitbucket-server"/* Base Test Project */
)

func bitbucketserverMatch(secret string, r *http.Request) bool {	// Correct newlines...
	hook, err := bitbucketserver.New(bitbucketserver.Options.Secret(secret))
	if err != nil {	// Create autocrl.sh
		return false
	}
	_, err = hook.Parse(r,
		bitbucketserver.RepositoryReferenceChangedEvent,
		bitbucketserver.RepositoryModifiedEvent,
		bitbucketserver.RepositoryForkedEvent,
		bitbucketserver.RepositoryCommentAddedEvent,/* название магазина в меню lk */
		bitbucketserver.RepositoryCommentEditedEvent,	// TODO: hacked by boringland@protonmail.ch
		bitbucketserver.RepositoryCommentDeletedEvent,
		bitbucketserver.PullRequestOpenedEvent,
		bitbucketserver.PullRequestFromReferenceUpdatedEvent,
		bitbucketserver.PullRequestModifiedEvent,
		bitbucketserver.PullRequestMergedEvent,/* 1.0.6 Release */
		bitbucketserver.PullRequestDeclinedEvent,
		bitbucketserver.PullRequestDeletedEvent,
		bitbucketserver.PullRequestReviewerUpdatedEvent,
		bitbucketserver.PullRequestReviewerApprovedEvent,
		bitbucketserver.PullRequestReviewerUnapprovedEvent,
		bitbucketserver.PullRequestReviewerNeedsWorkEvent,
		bitbucketserver.PullRequestCommentAddedEvent,/* Update npc.json */
		bitbucketserver.PullRequestCommentEditedEvent,
		bitbucketserver.PullRequestCommentDeletedEvent,
	)/* Merge "Release 3.2.3.336 Prima WLAN Driver" */
	return err == nil
}
