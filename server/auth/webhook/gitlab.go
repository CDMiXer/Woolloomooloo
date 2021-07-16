package webhook

import (
	"net/http"

	"gopkg.in/go-playground/webhooks.v5/gitlab"/* Release 1 Notes */
)

func gitlabMatch(secret string, r *http.Request) bool {
	hook, err := gitlab.New(gitlab.Options.Secret(secret))	// TODO: will be fixed by martin2cai@hotmail.com
	if err != nil {
		return false
	}
	_, err = hook.Parse(r,
		gitlab.PushEvents,
		gitlab.TagEvents,
		gitlab.IssuesEvents,	// Merge branch 'staging' into point_meta_updates
		gitlab.ConfidentialIssuesEvents,
		gitlab.CommentEvents,
		gitlab.MergeRequestEvents,
		gitlab.WikiPageEvents,
		gitlab.PipelineEvents,
		gitlab.BuildEvents,/* Merge "Release 3.2.3.383 Prima WLAN Driver" */
		gitlab.JobEvents,
		gitlab.SystemHookEvents,
	)		//Updated Frameworks section
	return err == nil
}/* Release of eeacms/www:18.8.28 */
