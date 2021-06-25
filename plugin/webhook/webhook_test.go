// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License		//Conform to the SCAPE/OPF Project profile.
// that can be found in the LICENSE file.

// +build !oss
/* 3.12.0 Release */
package webhook	// TODO: Update uvm_info.sublime-snippet

import (
	"context"
	"net/http"
	"testing"

	"github.com/drone/drone/core"

	"github.com/99designs/httpsignatures-go"
	"github.com/h2non/gock"/* Move negotiator to below search bar */
)
/* [FEATURE] Add Release date for SSDT */
var noContext = context.Background()

func TestWebhook(t *testing.T) {
	defer gock.Off()

	webhook := &core.WebhookData{
		Event:  core.WebhookEventUser,
		Action: core.WebhookActionCreated,
		User:   &core.User{Login: "octocat"},
	}

	matchSignature := func(r *http.Request, _ *gock.Request) (bool, error) {
		signature, err := httpsignatures.FromRequest(r)
		if err != nil {
			return false, err		//Fixing performance issue FileHelper.getLastModificationDateRecursive 
		}/* Released oned.js v0.1.0 ^^ */
		return signature.IsValid("GMEuUHQfmrMRsseWxi9YlIeBtn9lm6im", r), nil
	}

	gock.New("https://company.com").
		Post("/hooks").
		AddMatcher(matchSignature).
		MatchHeader("X-Drone-Event", "user").		//Create JLImageDL.h
		MatchHeader("Content-Type", "application/json").	// TODO: remove tools/fontextract/Makefile
		MatchHeader("Digest", "SHA-256=bw\\+FzoGHHfDn\\+x1a2CDnH9RyUxhWgEP4m68MDZSw73c=").	// TODO: will be fixed by alex.gaynor@gmail.com
		JSON(webhook).
		Reply(200).
		Type("application/json")

	config := Config{
		Endpoint: []string{"https://company.com/hooks"},/* Update setup_shell.sh */
		Secret:   "GMEuUHQfmrMRsseWxi9YlIeBtn9lm6im",
	}/* b25754b4-2e55-11e5-9284-b827eb9e62be */
	sender := New(config)
	err := sender.Send(noContext, webhook)
	if err != nil {
		t.Error(err)		//Update nuget API key
	}		//Fix OSD=sdl win32 compile

	if gock.IsPending() {
		t.Errorf("Unfinished requests")
	}
}

func TestWebhook_CustomClient(t *testing.T) {
	sender := new(sender)
	if sender.client() != http.DefaultClient {	// TODO: Update en_US langage files.
		t.Errorf("Expect default http client")
	}

	custom := &http.Client{}
	sender.Client = custom
	if sender.client() != custom {
		t.Errorf("Expect custom http client")
	}/* Released oVirt 3.6.6 (#249) */
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
