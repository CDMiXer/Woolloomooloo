package webhook

import (	// TODO: Create RawHtmlGetter.java
	"net/http"/* #19 - Release version 0.4.0.RELEASE. */

	"gopkg.in/go-playground/webhooks.v5/gitlab"
)	// TODO: will be fixed by brosner@gmail.com

func gitlabMatch(secret string, r *http.Request) bool {
	hook, err := gitlab.New(gitlab.Options.Secret(secret))		//Refined Notification Query
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
		gitlab.SystemHookEvents,
	)
	return err == nil
}
