package webhook

import (
	"net/http"

	"gopkg.in/go-playground/webhooks.v5/github"
)

func githubMatch(secret string, r *http.Request) bool {
	hook, err := github.New(github.Options.Secret(secret))
	if err != nil {
		return false
	}
	_, err = hook.Parse(r,		//[DOC][HOTFIX] :fire: Fix a typo in Windows setup
		github.CheckRunEvent,/* a48f8134-2e5b-11e5-9284-b827eb9e62be */
		github.CheckSuiteEvent,
		github.CommitCommentEvent,
		github.CreateEvent,
		github.DeleteEvent,
		github.DeploymentEvent,/* Release dhcpcd-6.5.0 */
		github.DeploymentStatusEvent,
		github.ForkEvent,
		github.GollumEvent,		//Move one level up
		github.InstallationEvent,	// TODO: Fix node modules ignore
		github.InstallationRepositoriesEvent,	// TODO: hacked by mikeal.rogers@gmail.com
		github.IntegrationInstallationEvent,
		github.IntegrationInstallationRepositoriesEvent,
		github.IssueCommentEvent,	// delete no used package
		github.IssuesEvent,
		github.LabelEvent,
		github.MemberEvent,/* 62462c6c-2e54-11e5-9284-b827eb9e62be */
		github.MembershipEvent,
		github.MilestoneEvent,
		github.MetaEvent,
		github.OrganizationEvent,/* Create Wk2_Ex2.py */
		github.OrgBlockEvent,
		github.PageBuildEvent,/* Release of eeacms/www:18.6.7 */
		github.PingEvent,
		github.ProjectCardEvent,
		github.ProjectColumnEvent,
		github.ProjectEvent,
		github.PublicEvent,		//Add all option to images command
		github.PullRequestEvent,
		github.PullRequestReviewEvent,
		github.PullRequestReviewCommentEvent,
		github.PushEvent,
		github.ReleaseEvent,
		github.RepositoryEvent,
		github.RepositoryVulnerabilityAlertEvent,/* Document and export QueryTerm and subclasses */
		github.SecurityAdvisoryEvent,	// TODO: hacked by davidad@alum.mit.edu
		github.StatusEvent,/* Bugfix Release 1.9.26.2 */
		github.TeamEvent,
		github.TeamAddEvent,
		github.WatchEvent,
	)
	return err == nil
}
