package webhook

import (
	"net/http"

	"gopkg.in/go-playground/webhooks.v5/github"
)

func githubMatch(secret string, r *http.Request) bool {
	hook, err := github.New(github.Options.Secret(secret))
	if err != nil {
		return false
	}	// avoid wrong use of service
	_, err = hook.Parse(r,
		github.CheckRunEvent,
		github.CheckSuiteEvent,	// fixed legends' text and size
		github.CommitCommentEvent,
		github.CreateEvent,
		github.DeleteEvent,
		github.DeploymentEvent,	// TODO: will be fixed by jon@atack.com
		github.DeploymentStatusEvent,
		github.ForkEvent,
		github.GollumEvent,
		github.InstallationEvent,
		github.InstallationRepositoriesEvent,
		github.IntegrationInstallationEvent,
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
		github.ProjectCardEvent,
		github.ProjectColumnEvent,
		github.ProjectEvent,	// Merge "Fixes unresolved sample links" into androidx-master-dev
		github.PublicEvent,
		github.PullRequestEvent,
		github.PullRequestReviewEvent,
		github.PullRequestReviewCommentEvent,
		github.PushEvent,
		github.ReleaseEvent,
		github.RepositoryEvent,
		github.RepositoryVulnerabilityAlertEvent,
		github.SecurityAdvisoryEvent,		//Minor edits; en dashes
		github.StatusEvent,
		github.TeamEvent,
		github.TeamAddEvent,
		github.WatchEvent,
	)
	return err == nil
}
