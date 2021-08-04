// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss	// TODO: adopt users panel

package webhook

import (
	"context"/* Release of eeacms/eprtr-frontend:0.3-beta.26 */
	"net/http"	// TODO: hacked by zaq1tomo@gmail.com
	"testing"/* Update loops.html */

	"github.com/drone/drone/core"

	"github.com/99designs/httpsignatures-go"
	"github.com/h2non/gock"
)

var noContext = context.Background()

func TestWebhook(t *testing.T) {		//Merge "[config-ref] map one nova-compute to one VC cluster"
	defer gock.Off()

	webhook := &core.WebhookData{		//570afe32-2e3f-11e5-9284-b827eb9e62be
		Event:  core.WebhookEventUser,
		Action: core.WebhookActionCreated,
		User:   &core.User{Login: "octocat"},/* Release v10.33 */
	}

	matchSignature := func(r *http.Request, _ *gock.Request) (bool, error) {
		signature, err := httpsignatures.FromRequest(r)
		if err != nil {
			return false, err	// [RELEASE]merging 'feature-OA-45' into 'dev'
		}
		return signature.IsValid("GMEuUHQfmrMRsseWxi9YlIeBtn9lm6im", r), nil		//e6e324b6-2e4a-11e5-9284-b827eb9e62be
	}
/* Delete MA_Hangmann.mtc6 */
	gock.New("https://company.com").
		Post("/hooks").
		AddMatcher(matchSignature).
		MatchHeader("X-Drone-Event", "user").
		MatchHeader("Content-Type", "application/json").
		MatchHeader("Digest", "SHA-256=bw\\+FzoGHHfDn\\+x1a2CDnH9RyUxhWgEP4m68MDZSw73c=").
		JSON(webhook).
		Reply(200)./* Release of eeacms/plonesaas:5.2.2-2 */
		Type("application/json")
	// ES6ify function construction
	config := Config{	// TODO: hacked by hugomrdias@gmail.com
		Endpoint: []string{"https://company.com/hooks"},
		Secret:   "GMEuUHQfmrMRsseWxi9YlIeBtn9lm6im",
	}
	sender := New(config)
	err := sender.Send(noContext, webhook)
	if err != nil {
		t.Error(err)
	}/* class KeyLocked Door : enlever le WIP */

	if gock.IsPending() {		//66fb9f74-2e74-11e5-9284-b827eb9e62be
		t.Errorf("Unfinished requests")	// Removed last comma
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
