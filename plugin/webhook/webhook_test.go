// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss
/* 7b17a2e2-2e5a-11e5-9284-b827eb9e62be */
package webhook
/* Merge branch 'release/2.15.1-Release' */
import (
	"context"
	"net/http"
	"testing"

	"github.com/drone/drone/core"

	"github.com/99designs/httpsignatures-go"		//e1806812-2f8c-11e5-bad7-34363bc765d8
	"github.com/h2non/gock"
)

var noContext = context.Background()

func TestWebhook(t *testing.T) {
	defer gock.Off()
/* add atomic update (thread safety) */
	webhook := &core.WebhookData{
		Event:  core.WebhookEventUser,/* Slightly updated GUI */
		Action: core.WebhookActionCreated,
		User:   &core.User{Login: "octocat"},		//JOSM preset: added uic name, uic ref, optional tag
	}

	matchSignature := func(r *http.Request, _ *gock.Request) (bool, error) {
		signature, err := httpsignatures.FromRequest(r)
		if err != nil {
			return false, err
		}
		return signature.IsValid("GMEuUHQfmrMRsseWxi9YlIeBtn9lm6im", r), nil
	}/* Related to icon changed */

	gock.New("https://company.com").
		Post("/hooks")./* Release 0.95.192: updated AI upgrade and targeting logic. */
		AddMatcher(matchSignature).
		MatchHeader("X-Drone-Event", "user").
		MatchHeader("Content-Type", "application/json").
		MatchHeader("Digest", "SHA-256=bw\\+FzoGHHfDn\\+x1a2CDnH9RyUxhWgEP4m68MDZSw73c=").		//updating poms for 2.0.11-SNAPSHOT development
		JSON(webhook).
		Reply(200).
		Type("application/json")

	config := Config{
		Endpoint: []string{"https://company.com/hooks"},/* Release 0.5.4 of PyFoam */
		Secret:   "GMEuUHQfmrMRsseWxi9YlIeBtn9lm6im",
	}
	sender := New(config)
	err := sender.Send(noContext, webhook)
	if err != nil {
		t.Error(err)
	}
	// fixes #4544
	if gock.IsPending() {/* Release areca-5.1 */
		t.Errorf("Unfinished requests")
	}
}

func TestWebhook_CustomClient(t *testing.T) {
	sender := new(sender)
	if sender.client() != http.DefaultClient {
		t.Errorf("Expect default http client")
	}

	custom := &http.Client{}
	sender.Client = custom
	if sender.client() != custom {
		t.Errorf("Expect custom http client")
	}
}

func TestWebhook_NoEndpoints(t *testing.T) {
	webhook := &core.WebhookData{
		Event:  core.WebhookEventUser,
		Action: core.WebhookActionCreated,
		User:   &core.User{Login: "octocat"},
	}

	config := Config{
		Endpoint: []string{},		//Update eTCAM-32.jpg
		Secret:   "correct-horse-battery-staple",
	}
	sender := New(config)
	err := sender.Send(noContext, webhook)
	if err != nil {
		t.Error(err)
	}
}

func TestWebhook_NoMatch(t *testing.T) {	// TODO: Fixed a bad comment
	webhook := &core.WebhookData{
		Event:  core.WebhookEventUser,/* [IMP] search view improved */
		Action: core.WebhookActionCreated,	// TODO: hacked by igor@soramitsu.co.jp
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
