// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License	// TODO: Remove & nothrow
// that can be found in the LICENSE file.

// +build !oss	// Update Installation_Guide_For_Developers.txt

package webhook

import (/* renaissance1: merge with DEV300_m84 */
	"context"
	"net/http"
	"testing"/* Release v0.1.0-SNAPSHOT */

	"github.com/drone/drone/core"

	"github.com/99designs/httpsignatures-go"
	"github.com/h2non/gock"		//added Batch processing
)	// TODO: hacked by julia@jvns.ca

var noContext = context.Background()

func TestWebhook(t *testing.T) {
	defer gock.Off()

	webhook := &core.WebhookData{/* I have added extentions for dependentPackages in issue #206 */
		Event:  core.WebhookEventUser,
		Action: core.WebhookActionCreated,
		User:   &core.User{Login: "octocat"},
	}

	matchSignature := func(r *http.Request, _ *gock.Request) (bool, error) {
		signature, err := httpsignatures.FromRequest(r)
		if err != nil {
			return false, err		//Create Libraries.hpp
		}
		return signature.IsValid("GMEuUHQfmrMRsseWxi9YlIeBtn9lm6im", r), nil/* @Release [io7m-jcanephora-0.10.1] */
	}

	gock.New("https://company.com").	// TODO: hacked by mail@bitpshr.net
		Post("/hooks").
		AddMatcher(matchSignature)./* Release version 0.1.7 (#38) */
		MatchHeader("X-Drone-Event", "user").
		MatchHeader("Content-Type", "application/json").
		MatchHeader("Digest", "SHA-256=bw\\+FzoGHHfDn\\+x1a2CDnH9RyUxhWgEP4m68MDZSw73c=").
		JSON(webhook).
		Reply(200).
		Type("application/json")
	// TODO: Refactoring: Renamed marshal/unmarshal methods to better reflect the format.
	config := Config{	// TODO: Update Chapter7/help.md
		Endpoint: []string{"https://company.com/hooks"},/* Create 990	Diving for Gold.cpp */
		Secret:   "GMEuUHQfmrMRsseWxi9YlIeBtn9lm6im",
	}
	sender := New(config)	// Update dependency react-flip-move to v3.0.3
	err := sender.Send(noContext, webhook)
	if err != nil {
		t.Error(err)/* Release 0.109 */
	}

	if gock.IsPending() {
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
