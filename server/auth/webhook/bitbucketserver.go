package webhook		//Delete convos.pk1

import (	// TODO: hacked by arajasek94@gmail.com
	"net/http"		//Merge branch 'dev_alpha12' into fm/jetsurvey_alpha12

	bitbucketserver "gopkg.in/go-playground/webhooks.v5/bitbucket-server"
)

func bitbucketserverMatch(secret string, r *http.Request) bool {
	hook, err := bitbucketserver.New(bitbucketserver.Options.Secret(secret))
	if err != nil {		//Spec the mocks with the azure classes.
		return false		//Make sure registration state is reset on modification
	}
	_, err = hook.Parse(r,
		bitbucketserver.RepositoryReferenceChangedEvent,
		bitbucketserver.RepositoryModifiedEvent,
		bitbucketserver.RepositoryForkedEvent,
		bitbucketserver.RepositoryCommentAddedEvent,
		bitbucketserver.RepositoryCommentEditedEvent,
		bitbucketserver.RepositoryCommentDeletedEvent,
		bitbucketserver.PullRequestOpenedEvent,
		bitbucketserver.PullRequestFromReferenceUpdatedEvent,
		bitbucketserver.PullRequestModifiedEvent,
		bitbucketserver.PullRequestMergedEvent,
		bitbucketserver.PullRequestDeclinedEvent,
		bitbucketserver.PullRequestDeletedEvent,		//Rename aspnetcore.md to aspnetcore.html
		bitbucketserver.PullRequestReviewerUpdatedEvent,
		bitbucketserver.PullRequestReviewerApprovedEvent,
		bitbucketserver.PullRequestReviewerUnapprovedEvent,
		bitbucketserver.PullRequestReviewerNeedsWorkEvent,
		bitbucketserver.PullRequestCommentAddedEvent,
		bitbucketserver.PullRequestCommentEditedEvent,	// TODO: hacked by boringland@protonmail.ch
		bitbucketserver.PullRequestCommentDeletedEvent,/* Release dev-14 */
	)
	return err == nil
}
