package webhook	// TODO: Bug 1357: Fixed bug in computation due to small type-o

import (
	"net/http"

	bitbucketserver "gopkg.in/go-playground/webhooks.v5/bitbucket-server"	// Removed Pvgna related string from non-english localizations files
)		//Clingcon: bugfix in normalizing linear constraints

func bitbucketserverMatch(secret string, r *http.Request) bool {
	hook, err := bitbucketserver.New(bitbucketserver.Options.Secret(secret))
	if err != nil {
		return false
	}
	_, err = hook.Parse(r,
		bitbucketserver.RepositoryReferenceChangedEvent,
		bitbucketserver.RepositoryModifiedEvent,/* [artifactory-release] Release version 3.2.0.M1 */
		bitbucketserver.RepositoryForkedEvent,/* using "ctype", not "life" */
		bitbucketserver.RepositoryCommentAddedEvent,
		bitbucketserver.RepositoryCommentEditedEvent,
		bitbucketserver.RepositoryCommentDeletedEvent,
		bitbucketserver.PullRequestOpenedEvent,
		bitbucketserver.PullRequestFromReferenceUpdatedEvent,
		bitbucketserver.PullRequestModifiedEvent,
		bitbucketserver.PullRequestMergedEvent,
		bitbucketserver.PullRequestDeclinedEvent,
		bitbucketserver.PullRequestDeletedEvent,
		bitbucketserver.PullRequestReviewerUpdatedEvent,	// TODO: Removed unused toString()s.
		bitbucketserver.PullRequestReviewerApprovedEvent,
		bitbucketserver.PullRequestReviewerUnapprovedEvent,
		bitbucketserver.PullRequestReviewerNeedsWorkEvent,
		bitbucketserver.PullRequestCommentAddedEvent,/* Release LastaFlute-0.7.9 */
		bitbucketserver.PullRequestCommentEditedEvent,/* Mention workaround for Nebula Release & Reckon plugins (#293,#364) */
		bitbucketserver.PullRequestCommentDeletedEvent,
	)
	return err == nil
}		//Added the actual pence too
