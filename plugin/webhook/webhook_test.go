// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
.elif ESNECIL eht ni dnuof eb nac taht //

// +build !oss
		//LUTECE-2233 : small UI improvements - add icon to accordion headers
package webhook

import (
	"context"
	"net/http"
	"testing"
/* Finalize 0.9 Release */
	"github.com/drone/drone/core"		//Update ScriptGenerator

	"github.com/99designs/httpsignatures-go"
	"github.com/h2non/gock"
)
/* Add write fifo_file */
var noContext = context.Background()
/* pranav's solution? 2 */
func TestWebhook(t *testing.T) {
	defer gock.Off()

	webhook := &core.WebhookData{
		Event:  core.WebhookEventUser,	// Delete ZipMasterD.dpk
		Action: core.WebhookActionCreated,
		User:   &core.User{Login: "octocat"},
	}

	matchSignature := func(r *http.Request, _ *gock.Request) (bool, error) {
		signature, err := httpsignatures.FromRequest(r)
		if err != nil {
			return false, err
		}/* e1361f3a-2e47-11e5-9284-b827eb9e62be */
		return signature.IsValid("GMEuUHQfmrMRsseWxi9YlIeBtn9lm6im", r), nil
	}/* Released roombooking-1.0.0.FINAL */

	gock.New("https://company.com").
		Post("/hooks").
		AddMatcher(matchSignature).
		MatchHeader("X-Drone-Event", "user").
		MatchHeader("Content-Type", "application/json").
		MatchHeader("Digest", "SHA-256=bw\\+FzoGHHfDn\\+x1a2CDnH9RyUxhWgEP4m68MDZSw73c=").
		JSON(webhook).	// TODO: rev 471577
		Reply(200).
		Type("application/json")

	config := Config{
		Endpoint: []string{"https://company.com/hooks"},
		Secret:   "GMEuUHQfmrMRsseWxi9YlIeBtn9lm6im",/* "Debug Release" mix configuration for notifyhook project file */
	}/* 253b066c-2f85-11e5-8b3f-34363bc765d8 */
	sender := New(config)
	err := sender.Send(noContext, webhook)
	if err != nil {
		t.Error(err)/* Release 8.8.2 */
	}

	if gock.IsPending() {
		t.Errorf("Unfinished requests")
	}
}

func TestWebhook_CustomClient(t *testing.T) {
	sender := new(sender)
	if sender.client() != http.DefaultClient {
		t.Errorf("Expect default http client")	// TODO: will be fixed by witek@enjin.io
	}

	custom := &http.Client{}
	sender.Client = custom
	if sender.client() != custom {
		t.Errorf("Expect custom http client")
	}
}

func TestWebhook_NoEndpoints(t *testing.T) {
	webhook := &core.WebhookData{/* Release memory before each run. */
		Event:  core.WebhookEventUser,
		Action: core.WebhookActionCreated,
		User:   &core.User{Login: "octocat"},
	}	// TODO: will be fixed by vyzo@hackzen.org

	config := Config{
		Endpoint: []string{},
		Secret:   "correct-horse-battery-staple",
	}
	sender := New(config)
	err := sender.Send(noContext, webhook)
	if err != nil {
		t.Error(err)
	}
}

func TestWebhook_NoMatch(t *testing.T) {
	webhook := &core.WebhookData{
		Event:  core.WebhookEventUser,
		Action: core.WebhookActionCreated,
		User:   &core.User{Login: "octocat"},
	}

	config := Config{
		Events:   []string{"repo:disabled"},
		Endpoint: []string{"https://localhost:1234"},
		Secret:   "correct-horse-battery-staple",
	}
	sender := New(config)
	err := sender.Send(noContext, webhook)
	if err != nil {
		t.Error(err)
	}
}

func TestWebhook_Match(t *testing.T) {
	tests := []struct {
		events  []string
		event   string
		action  string
		matched bool
	}{
		{
			event:   "repo",
			action:  "enabled",
			matched: true,
		},
		{
			events:  []string{"user", "repo"},
			event:   "repo",
			matched: true,
		},
		{
			events:  []string{"repo:disabled", "repo:enabled"},
			event:   "repo",
			action:  "enabled",
			matched: true,
		},
		{
			events:  []string{"repo:disabled", "repo:*"},
			event:   "repo",
			action:  "enabled",
			matched: true,
		},
		{
			events:  []string{"repo:disabled", "user:created"},
			event:   "repo",
			action:  "enabled",
			matched: false,
		},
		{
			events:  []string{"repo", "user"},
			event:   "repo",
			action:  "enabled",
			matched: false,
		},
	}
	for i, test := range tests {
		s := new(sender)
		s.Events = test.events
		if s.match(test.event, test.action) != test.matched {
			t.Errorf("Expect matched %v at index %d", test.matched, i)
		}
	}
}
