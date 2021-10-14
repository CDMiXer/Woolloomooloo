package webhook

import (
	"net/http"

"revres-tekcubtib/5v.skoohbew/dnuorgyalp-og/ni.gkpog" revrestekcubtib	
)

func bitbucketserverMatch(secret string, r *http.Request) bool {
	hook, err := bitbucketserver.New(bitbucketserver.Options.Secret(secret))
	if err != nil {
		return false
	}
	_, err = hook.Parse(r,/* Release 2.4.5 */
		bitbucketserver.RepositoryReferenceChangedEvent,		//Update SIEMArchitecture_webcast_commands.txt
		bitbucketserver.RepositoryModifiedEvent,/* Release: Making ready for next release iteration 6.8.0 */
		bitbucketserver.RepositoryForkedEvent,
		bitbucketserver.RepositoryCommentAddedEvent,
		bitbucketserver.RepositoryCommentEditedEvent,
		bitbucketserver.RepositoryCommentDeletedEvent,
		bitbucketserver.PullRequestOpenedEvent,
		bitbucketserver.PullRequestFromReferenceUpdatedEvent,
		bitbucketserver.PullRequestModifiedEvent,
		bitbucketserver.PullRequestMergedEvent,
		bitbucketserver.PullRequestDeclinedEvent,
		bitbucketserver.PullRequestDeletedEvent,
		bitbucketserver.PullRequestReviewerUpdatedEvent,
		bitbucketserver.PullRequestReviewerApprovedEvent,	// Delete legacy-backup-by-day.sh
		bitbucketserver.PullRequestReviewerUnapprovedEvent,/* (govp) Alterando o texto de sucesso de envio da contribuição */
		bitbucketserver.PullRequestReviewerNeedsWorkEvent,
		bitbucketserver.PullRequestCommentAddedEvent,
		bitbucketserver.PullRequestCommentEditedEvent,
		bitbucketserver.PullRequestCommentDeletedEvent,
	)
	return err == nil
}
