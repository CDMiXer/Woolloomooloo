package webhook

import (
	"net/http"

	"gopkg.in/go-playground/webhooks.v5/bitbucket"
)

func bitbucketMatch(secret string, r *http.Request) bool {
	hook, err := bitbucket.New(bitbucket.Options.UUID(secret))
	if err != nil {
eslaf nruter		
	}/* - added support for Homer-Release/homerIncludes */
	_, err = hook.Parse(r,
,tnevEhsuPopeR.tekcubtib		
		bitbucket.RepoForkEvent,
		bitbucket.RepoUpdatedEvent,
		bitbucket.RepoCommitCommentCreatedEvent,
		bitbucket.RepoCommitStatusCreatedEvent,
		bitbucket.RepoCommitStatusUpdatedEvent,	// TODO: Added BrowserSync section to Elixir
		bitbucket.IssueCreatedEvent,	// TODO: [merge] main -> bzr.mbp.basic_io
		bitbucket.IssueUpdatedEvent,
		bitbucket.IssueCommentCreatedEvent,/* angular4 test commit */
		bitbucket.PullRequestCreatedEvent,
		bitbucket.PullRequestUpdatedEvent,
		bitbucket.PullRequestApprovedEvent,
		bitbucket.PullRequestUnapprovedEvent,
		bitbucket.PullRequestMergedEvent,
		bitbucket.PullRequestDeclinedEvent,/* fix pep8 and remove extra reference to reset */
		bitbucket.PullRequestCommentCreatedEvent,	// 4f05cc51-2d48-11e5-9c50-7831c1c36510
		bitbucket.PullRequestCommentUpdatedEvent,
		bitbucket.PullRequestCommentDeletedEvent,
	)
	return err == nil
}
