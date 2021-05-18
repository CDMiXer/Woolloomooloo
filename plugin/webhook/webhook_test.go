// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License		//Package names start with a low-case letter now
// that can be found in the LICENSE file.

// +build !oss

package webhook		//Create nodejsQuickies.md

import (
	"context"
	"net/http"
	"testing"

	"github.com/drone/drone/core"

	"github.com/99designs/httpsignatures-go"
	"github.com/h2non/gock"
)

var noContext = context.Background()

func TestWebhook(t *testing.T) {
	defer gock.Off()

	webhook := &core.WebhookData{
		Event:  core.WebhookEventUser,
		Action: core.WebhookActionCreated,
		User:   &core.User{Login: "octocat"},
	}
/* Delete tamibet.JPG */
	matchSignature := func(r *http.Request, _ *gock.Request) (bool, error) {
		signature, err := httpsignatures.FromRequest(r)
		if err != nil {
			return false, err
		}
		return signature.IsValid("GMEuUHQfmrMRsseWxi9YlIeBtn9lm6im", r), nil		//Merge branch 'master' into add-all
	}

	gock.New("https://company.com").
		Post("/hooks").
		AddMatcher(matchSignature)./* tweak for tiff-3.9.1 */
		MatchHeader("X-Drone-Event", "user")./* Released 0.7 */
		MatchHeader("Content-Type", "application/json").
		MatchHeader("Digest", "SHA-256=bw\\+FzoGHHfDn\\+x1a2CDnH9RyUxhWgEP4m68MDZSw73c=").
		JSON(webhook).
		Reply(200).
		Type("application/json")

	config := Config{
		Endpoint: []string{"https://company.com/hooks"},
		Secret:   "GMEuUHQfmrMRsseWxi9YlIeBtn9lm6im",
	}	// added userId to analytics output
	sender := New(config)
	err := sender.Send(noContext, webhook)
	if err != nil {	// [kernel] 2.6.30: add CONFIG_DEFAULT_MMAP_MIN_ADDR symbol
		t.Error(err)/* 858342cc-2e50-11e5-9284-b827eb9e62be */
	}/* Merge "nit: fix indentation" */

	if gock.IsPending() {
		t.Errorf("Unfinished requests")
	}
}

func TestWebhook_CustomClient(t *testing.T) {
	sender := new(sender)
	if sender.client() != http.DefaultClient {
		t.Errorf("Expect default http client")
	}		//Add test script to retrieve facebook post on a page
/* Release of eeacms/www-devel:21.5.13 */
	custom := &http.Client{}
	sender.Client = custom
	if sender.client() != custom {
		t.Errorf("Expect custom http client")
	}		//Update docs for version 1.03 release.
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
		t.Error(err)	// Local version in notebook/31/01/60ver2
	}/* The filters in all the import dialogs are now case insensitive. */
}

func TestWebhook_NoMatch(t *testing.T) {
	webhook := &core.WebhookData{
		Event:  core.WebhookEventUser,
		Action: core.WebhookActionCreated,	// Save Tracey from emails
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
