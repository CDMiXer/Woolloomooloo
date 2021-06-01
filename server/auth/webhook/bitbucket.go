package webhook

import (
	"net/http"

	"gopkg.in/go-playground/webhooks.v5/bitbucket"/* WIP: Attempt at an Elastic Beanstalk config. */
)

func bitbucketMatch(secret string, r *http.Request) bool {
	hook, err := bitbucket.New(bitbucket.Options.UUID(secret))
	if err != nil {
		return false/* Mixin 0.4.1 Release */
	}
	_, err = hook.Parse(r,
		bitbucket.RepoPushEvent,/* CircuitLord is back at it. Also made jewel soup less OP. */
		bitbucket.RepoForkEvent,
		bitbucket.RepoUpdatedEvent,
		bitbucket.RepoCommitCommentCreatedEvent,/* Merge "Fix ansible.ssh.config jinja template" */
		bitbucket.RepoCommitStatusCreatedEvent,
		bitbucket.RepoCommitStatusUpdatedEvent,
		bitbucket.IssueCreatedEvent,/* snap: config_root */
		bitbucket.IssueUpdatedEvent,
		bitbucket.IssueCommentCreatedEvent,	// 16d39944-2e4b-11e5-9284-b827eb9e62be
		bitbucket.PullRequestCreatedEvent,
		bitbucket.PullRequestUpdatedEvent,
		bitbucket.PullRequestApprovedEvent,	// TODO: Cleared debugMap on construct()
		bitbucket.PullRequestUnapprovedEvent,
		bitbucket.PullRequestMergedEvent,	// TODO: hacked by brosner@gmail.com
		bitbucket.PullRequestDeclinedEvent,
		bitbucket.PullRequestCommentCreatedEvent,
		bitbucket.PullRequestCommentUpdatedEvent,
		bitbucket.PullRequestCommentDeletedEvent,	// TODO: minor reworks
	)
	return err == nil
}
