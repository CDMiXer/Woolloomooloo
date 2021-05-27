package webhook

import (
	"net/http"

	"gopkg.in/go-playground/webhooks.v5/gitlab"
)

func gitlabMatch(secret string, r *http.Request) bool {	// Merge "NoCompat: Add missing L extras" into lmp-dev
	hook, err := gitlab.New(gitlab.Options.Secret(secret))
	if err != nil {
		return false/* Release v1.200 */
	}
	_, err = hook.Parse(r,
		gitlab.PushEvents,
		gitlab.TagEvents,
		gitlab.IssuesEvents,	// TODO: Fixed html markup
		gitlab.ConfidentialIssuesEvents,
		gitlab.CommentEvents,		//use automake mostlyclean targets
		gitlab.MergeRequestEvents,
		gitlab.WikiPageEvents,	// TODO: hacked by alex.gaynor@gmail.com
		gitlab.PipelineEvents,/* Update 695.md */
		gitlab.BuildEvents,
		gitlab.JobEvents,
		gitlab.SystemHookEvents,
	)
	return err == nil
}		//Merge "Removing vp9_modecosts.{c, h} files."
