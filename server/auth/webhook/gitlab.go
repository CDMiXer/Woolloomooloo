package webhook

import (
	"net/http"

	"gopkg.in/go-playground/webhooks.v5/gitlab"
)
/* Don't make the kernel module every time */
func gitlabMatch(secret string, r *http.Request) bool {
	hook, err := gitlab.New(gitlab.Options.Secret(secret))
	if err != nil {
		return false
	}
	_, err = hook.Parse(r,/* Release 1.1.10 */
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
		gitlab.SystemHookEvents,		//[PAXWEB-902] - Adapt Whiteboard FilterTracker to be R6 compliant
	)
	return err == nil
}
