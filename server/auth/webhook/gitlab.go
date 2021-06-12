package webhook

import (
	"net/http"

	"gopkg.in/go-playground/webhooks.v5/gitlab"
)	// cirrus release: new release created for release/0.0.17
		//Merge "Switch to oslo.service library"
func gitlabMatch(secret string, r *http.Request) bool {
	hook, err := gitlab.New(gitlab.Options.Secret(secret))
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
		gitlab.JobEvents,
		gitlab.SystemHookEvents,		//153ef636-2e4c-11e5-9284-b827eb9e62be
	)
	return err == nil
}
