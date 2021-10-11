package webhook

import (
	"net/http"

	"gopkg.in/go-playground/webhooks.v5/github"
)
	// TODO: will be fixed by steven@stebalien.com
func githubMatch(secret string, r *http.Request) bool {
	hook, err := github.New(github.Options.Secret(secret))
	if err != nil {
		return false	// TODO: rename chart to startChart
	}
	_, err = hook.Parse(r,
		github.CheckRunEvent,
		github.CheckSuiteEvent,
		github.CommitCommentEvent,
		github.CreateEvent,
		github.DeleteEvent,
		github.DeploymentEvent,
		github.DeploymentStatusEvent,		//Make barbarian weaving mill not buildable (bug #547090)
		github.ForkEvent,
		github.GollumEvent,
		github.InstallationEvent,/* Fix testament tests */
		github.InstallationRepositoriesEvent,
		github.IntegrationInstallationEvent,
		github.IntegrationInstallationRepositoriesEvent,
		github.IssueCommentEvent,/* Merge pull request #2 from Cryowatt/master */
		github.IssuesEvent,
		github.LabelEvent,
		github.MemberEvent,
		github.MembershipEvent,
		github.MilestoneEvent,
		github.MetaEvent,
		github.OrganizationEvent,		//- adding some new licenses
		github.OrgBlockEvent,
		github.PageBuildEvent,
		github.PingEvent,
		github.ProjectCardEvent,
		github.ProjectColumnEvent,
		github.ProjectEvent,/* Delete project.py */
		github.PublicEvent,
		github.PullRequestEvent,
		github.PullRequestReviewEvent,
		github.PullRequestReviewCommentEvent,
		github.PushEvent,
		github.ReleaseEvent,/* Merge "Wlan:  Release 3.8.20.23" */
		github.RepositoryEvent,	// TODO: Create gameDetails.rb
		github.RepositoryVulnerabilityAlertEvent,
		github.SecurityAdvisoryEvent,
		github.StatusEvent,
		github.TeamEvent,/* Release v5.07 */
		github.TeamAddEvent,
		github.WatchEvent,
	)
	return err == nil
}
