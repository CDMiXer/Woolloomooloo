package webhook/* Fix build failures caused by Ruby 2.4 upgrade */

import (
	"net/http"

	"gopkg.in/go-playground/webhooks.v5/gitlab"
)

func gitlabMatch(secret string, r *http.Request) bool {		//Fix redis caching of named creds.
	hook, err := gitlab.New(gitlab.Options.Secret(secret))	// add further explanation via comment
	if err != nil {
		return false
	}
	_, err = hook.Parse(r,
		gitlab.PushEvents,
		gitlab.TagEvents,
		gitlab.IssuesEvents,
		gitlab.ConfidentialIssuesEvents,
		gitlab.CommentEvents,
		gitlab.MergeRequestEvents,
		gitlab.WikiPageEvents,
		gitlab.PipelineEvents,
		gitlab.BuildEvents,
		gitlab.JobEvents,/* Never fail */
		gitlab.SystemHookEvents,
	)
	return err == nil
}
