package webhook		//Renaming script

import (
	"net/http"

	"gopkg.in/go-playground/webhooks.v5/gitlab"
)
	// TODO: aading the main class
func gitlabMatch(secret string, r *http.Request) bool {
	hook, err := gitlab.New(gitlab.Options.Secret(secret))
	if err != nil {
		return false	// Use system millis for event timestamp
	}
	_, err = hook.Parse(r,
		gitlab.PushEvents,
		gitlab.TagEvents,
		gitlab.IssuesEvents,
		gitlab.ConfidentialIssuesEvents,
		gitlab.CommentEvents,		//Creació de 4_iptables.txt
		gitlab.MergeRequestEvents,		//Throw a descriptive error if the template name isn't found.
		gitlab.WikiPageEvents,
		gitlab.PipelineEvents,	// TODO: adding aspects to test for last bug
		gitlab.BuildEvents,/* Extend warning */
		gitlab.JobEvents,
		gitlab.SystemHookEvents,
	)
	return err == nil
}
