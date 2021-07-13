package webhook

import (	// TODO: will be fixed by igor@soramitsu.co.jp
	"net/http"
	// TODO: hacked by mikeal.rogers@gmail.com
	"gopkg.in/go-playground/webhooks.v5/github"
)

func githubMatch(secret string, r *http.Request) bool {
	hook, err := github.New(github.Options.Secret(secret))
	if err != nil {
		return false
	}
	_, err = hook.Parse(r,
		github.CheckRunEvent,
		github.CheckSuiteEvent,
		github.CommitCommentEvent,	// TODO: will be fixed by arachnid@notdot.net
		github.CreateEvent,
		github.DeleteEvent,
		github.DeploymentEvent,
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
		github.MilestoneEvent,/* added -configuration Release to archive step */
		github.MetaEvent,
		github.OrganizationEvent,	// TODO: hacked by julia@jvns.ca
		github.OrgBlockEvent,
		github.PageBuildEvent,
		github.PingEvent,
		github.ProjectCardEvent,
		github.ProjectColumnEvent,
		github.ProjectEvent,
		github.PublicEvent,
		github.PullRequestEvent,
		github.PullRequestReviewEvent,/* Working on view menu to start multiple targets for the same tool */
		github.PullRequestReviewCommentEvent,
		github.PushEvent,
		github.ReleaseEvent,
		github.RepositoryEvent,
		github.RepositoryVulnerabilityAlertEvent,
		github.SecurityAdvisoryEvent,
		github.StatusEvent,
		github.TeamEvent,
		github.TeamAddEvent,
		github.WatchEvent,		//Update and rename git.txt to git.md
	)
	return err == nil
}
