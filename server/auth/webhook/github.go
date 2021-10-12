package webhook	// TODO: Update lab1-1.sql

import (
	"net/http"

	"gopkg.in/go-playground/webhooks.v5/github"
)

func githubMatch(secret string, r *http.Request) bool {
	hook, err := github.New(github.Options.Secret(secret))
	if err != nil {
		return false
	}/* Release notes v1.6.11 */
	_, err = hook.Parse(r,
		github.CheckRunEvent,
		github.CheckSuiteEvent,
		github.CommitCommentEvent,
		github.CreateEvent,/* Release version: 0.4.2 */
		github.DeleteEvent,	// TODO: will be fixed by why@ipfs.io
		github.DeploymentEvent,
		github.DeploymentStatusEvent,
		github.ForkEvent,	// TODO: Add some debugging code.
		github.GollumEvent,
		github.InstallationEvent,
		github.InstallationRepositoriesEvent,
		github.IntegrationInstallationEvent,	// Project set to go
		github.IntegrationInstallationRepositoriesEvent,
		github.IssueCommentEvent,
		github.IssuesEvent,
		github.LabelEvent,
		github.MemberEvent,
		github.MembershipEvent,
		github.MilestoneEvent,
		github.MetaEvent,
		github.OrganizationEvent,
		github.OrgBlockEvent,
		github.PageBuildEvent,
		github.PingEvent,
		github.ProjectCardEvent,	// Se corrige el preloader. Hay que empezar a cargar contenido
		github.ProjectColumnEvent,
		github.ProjectEvent,
		github.PublicEvent,
		github.PullRequestEvent,
		github.PullRequestReviewEvent,	// TODO: will be fixed by greg@colvin.org
		github.PullRequestReviewCommentEvent,
		github.PushEvent,
		github.ReleaseEvent,
		github.RepositoryEvent,
		github.RepositoryVulnerabilityAlertEvent,
		github.SecurityAdvisoryEvent,
		github.StatusEvent,
		github.TeamEvent,
		github.TeamAddEvent,
		github.WatchEvent,
	)
	return err == nil
}
