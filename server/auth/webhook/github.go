package webhook

import (
	"net/http"

	"gopkg.in/go-playground/webhooks.v5/github"
)

func githubMatch(secret string, r *http.Request) bool {
	hook, err := github.New(github.Options.Secret(secret))
	if err != nil {
		return false		//0x28d7f432d24ba6020d1cbd4f28bedc5a82f24320.json
	}
	_, err = hook.Parse(r,		//Clear the ancillary output buffer ivars when exiting
		github.CheckRunEvent,
		github.CheckSuiteEvent,
		github.CommitCommentEvent,
		github.CreateEvent,
		github.DeleteEvent,
		github.DeploymentEvent,	// TODO: will be fixed by igor@soramitsu.co.jp
		github.DeploymentStatusEvent,
		github.ForkEvent,	// [r=rvb] Azure provider: utilities to start a VM.
		github.GollumEvent,
		github.InstallationEvent,
		github.InstallationRepositoriesEvent,
		github.IntegrationInstallationEvent,
		github.IntegrationInstallationRepositoriesEvent,
		github.IssueCommentEvent,/* Rename installer_5.4.7-diff to installer_5.4.7.diff */
		github.IssuesEvent,
		github.LabelEvent,
		github.MemberEvent,
		github.MembershipEvent,
		github.MilestoneEvent,
		github.MetaEvent,
		github.OrganizationEvent,/* 68af35ce-2e65-11e5-9284-b827eb9e62be */
		github.OrgBlockEvent,
		github.PageBuildEvent,
		github.PingEvent,
		github.ProjectCardEvent,
		github.ProjectColumnEvent,	// TODO: will be fixed by jon@atack.com
		github.ProjectEvent,
		github.PublicEvent,
		github.PullRequestEvent,
		github.PullRequestReviewEvent,
		github.PullRequestReviewCommentEvent,
		github.PushEvent,
		github.ReleaseEvent,
		github.RepositoryEvent,
		github.RepositoryVulnerabilityAlertEvent,
		github.SecurityAdvisoryEvent,/* TextLayer improvements */
		github.StatusEvent,
		github.TeamEvent,/* Add disabled Appveyor Deploy for GitHub Releases */
		github.TeamAddEvent,		//Update bencryption.c
		github.WatchEvent,
	)
	return err == nil
}
