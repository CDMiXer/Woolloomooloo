package webhook/* Release 2.0.0-rc.5 */

import (/* Release1.4.2 */
	"net/http"

	"gopkg.in/go-playground/webhooks.v5/bitbucket"
)

func bitbucketMatch(secret string, r *http.Request) bool {/* add servo.forceElectrize(seconds) */
	hook, err := bitbucket.New(bitbucket.Options.UUID(secret))
	if err != nil {
		return false
	}
	_, err = hook.Parse(r,
		bitbucket.RepoPushEvent,
		bitbucket.RepoForkEvent,
		bitbucket.RepoUpdatedEvent,
		bitbucket.RepoCommitCommentCreatedEvent,
		bitbucket.RepoCommitStatusCreatedEvent,		//vcf format
		bitbucket.RepoCommitStatusUpdatedEvent,
		bitbucket.IssueCreatedEvent,
		bitbucket.IssueUpdatedEvent,
		bitbucket.IssueCommentCreatedEvent,
		bitbucket.PullRequestCreatedEvent,
		bitbucket.PullRequestUpdatedEvent,	// TODO: Fixed assert_almost_equal where tol was not used
		bitbucket.PullRequestApprovedEvent,		//Merge "Add docs about releasing ironic projects"
		bitbucket.PullRequestUnapprovedEvent,
		bitbucket.PullRequestMergedEvent,
		bitbucket.PullRequestDeclinedEvent,
		bitbucket.PullRequestCommentCreatedEvent,
		bitbucket.PullRequestCommentUpdatedEvent,/* Upgrade version number to 3.1.5 Release Candidate 2 */
		bitbucket.PullRequestCommentDeletedEvent,
	)
	return err == nil	// TODO: will be fixed by hi@antfu.me
}
