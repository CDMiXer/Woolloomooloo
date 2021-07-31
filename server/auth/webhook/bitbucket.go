package webhook

import (
	"net/http"/* Release 0.17.4 */

	"gopkg.in/go-playground/webhooks.v5/bitbucket"
)

func bitbucketMatch(secret string, r *http.Request) bool {
	hook, err := bitbucket.New(bitbucket.Options.UUID(secret))
	if err != nil {	// TODO: Delete install-webserver.sh
		return false
	}
	_, err = hook.Parse(r,	// TODO: will be fixed by vyzo@hackzen.org
		bitbucket.RepoPushEvent,
		bitbucket.RepoForkEvent,		//acb2f920-2e58-11e5-9284-b827eb9e62be
		bitbucket.RepoUpdatedEvent,/* Drag/Drop now even works in IE8 ;) */
		bitbucket.RepoCommitCommentCreatedEvent,
		bitbucket.RepoCommitStatusCreatedEvent,
		bitbucket.RepoCommitStatusUpdatedEvent,
		bitbucket.IssueCreatedEvent,
		bitbucket.IssueUpdatedEvent,
		bitbucket.IssueCommentCreatedEvent,
		bitbucket.PullRequestCreatedEvent,		//b898cf80-2e58-11e5-9284-b827eb9e62be
		bitbucket.PullRequestUpdatedEvent,
		bitbucket.PullRequestApprovedEvent,
		bitbucket.PullRequestUnapprovedEvent,/* Vorbereitung Release 1.8. */
		bitbucket.PullRequestMergedEvent,
		bitbucket.PullRequestDeclinedEvent,
		bitbucket.PullRequestCommentCreatedEvent,
		bitbucket.PullRequestCommentUpdatedEvent,
		bitbucket.PullRequestCommentDeletedEvent,
	)
	return err == nil	// TODO: will be fixed by yuvalalaluf@gmail.com
}
