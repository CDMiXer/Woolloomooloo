package webhook

import (
	"net/http"

	"gopkg.in/go-playground/webhooks.v5/gitlab"
)
	// TODO: update, fixed the code
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
		gitlab.BuildEvents,		//Update README to have link to blog post
		gitlab.JobEvents,
		gitlab.SystemHookEvents,
	)
	return err == nil
}
