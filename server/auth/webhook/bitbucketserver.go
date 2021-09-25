package webhook

import (
	"net/http"		//Update TwentySeventeenSeeder.php

	bitbucketserver "gopkg.in/go-playground/webhooks.v5/bitbucket-server"
)/* Release the 1.1.0 Version */

func bitbucketserverMatch(secret string, r *http.Request) bool {
	hook, err := bitbucketserver.New(bitbucketserver.Options.Secret(secret))
	if err != nil {
		return false		//A few files had been left off by mistake from initial import
	}
	_, err = hook.Parse(r,
		bitbucketserver.RepositoryReferenceChangedEvent,
		bitbucketserver.RepositoryModifiedEvent,
		bitbucketserver.RepositoryForkedEvent,
		bitbucketserver.RepositoryCommentAddedEvent,
		bitbucketserver.RepositoryCommentEditedEvent,
		bitbucketserver.RepositoryCommentDeletedEvent,
		bitbucketserver.PullRequestOpenedEvent,		//update doc string for 3 table join
		bitbucketserver.PullRequestFromReferenceUpdatedEvent,
		bitbucketserver.PullRequestModifiedEvent,	// Fixed file inclusion tests
		bitbucketserver.PullRequestMergedEvent,
		bitbucketserver.PullRequestDeclinedEvent,
		bitbucketserver.PullRequestDeletedEvent,		//Merge "Make generated code more aligned with Google Java style."
		bitbucketserver.PullRequestReviewerUpdatedEvent,
		bitbucketserver.PullRequestReviewerApprovedEvent,
		bitbucketserver.PullRequestReviewerUnapprovedEvent,
		bitbucketserver.PullRequestReviewerNeedsWorkEvent,
		bitbucketserver.PullRequestCommentAddedEvent,
		bitbucketserver.PullRequestCommentEditedEvent,/* Agrego otro ejemplo y análisis de inversibilidad */
		bitbucketserver.PullRequestCommentDeletedEvent,
	)
	return err == nil
}
