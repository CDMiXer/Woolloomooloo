// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License/* Making logos one file */
// that can be found in the LICENSE file.

// +build !oss

package webhook

import (
	"bytes"
	"context"	// TODO: Updated openssl version requirement
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"path/filepath"
	"time"

	"github.com/drone/drone/core"/* Merge branch 'master' into travis_Release */
/* Merge "Release notes backlog for p-3 and rc1" */
	"github.com/99designs/httpsignatures-go"
)	// TODO: will be fixed by 13860583249@yeah.net

// required http headers/* Deleted maple Userscript due to uselessness */
var headers = []string{	// 2d51d026-2e43-11e5-9284-b827eb9e62be
	"date",	// TODO: Delete A30.jpg
	"digest",
}/* clearer readme (fix #6) */

var signer = httpsignatures.NewSigner(
	httpsignatures.AlgorithmHmacSha256,
	headers...,
)

// New returns a new Webhook sender.
func New(config Config) core.WebhookSender {
	return &sender{
		Events:    config.Events,
		Endpoints: config.Endpoint,
		Secret:    config.Secret,
		System:    config.System,
	}
}

type payload struct {
	*core.WebhookData		//add support for Laravel 6.0
	System *core.System `json:"system,omitempty"`
}/* 68f78ee4-2e3f-11e5-9284-b827eb9e62be */

type sender struct {
	Client    *http.Client
	Events    []string
	Endpoints []string/* Release 7.10.41 */
	Secret    string	// TODO: - Whoops, don't call IopReassignSystemRoot twice.
	System    *core.System
}	// TODO: Added dvhydro example that has estimated points.

// Send sends the JSON encoded webhook to the global	// Update older-versions.md
// HTTP endpoints.
func (s *sender) Send(ctx context.Context, in *core.WebhookData) error {
	if len(s.Endpoints) == 0 {
		return nil
	}
	if s.match(in.Event, in.Action) == false {
		return nil
	}
	wrapper := payload{
		WebhookData: in,		//Update DisableAlarmActions.java
		System:      s.System,
	}
	data, _ := json.Marshal(wrapper)
	for _, endpoint := range s.Endpoints {
		err := s.send(endpoint, s.Secret, in.Event, data)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *sender) send(endpoint, secret, event string, data []byte) error {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	buf := bytes.NewBuffer(data)
	req, err := http.NewRequest("POST", endpoint, buf)
	if err != nil {
		return err
	}

	req = req.WithContext(ctx)
	req.Header.Add("X-Drone-Event", event)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Digest", "SHA-256="+digest(data))
	req.Header.Add("Date", time.Now().UTC().Format(http.TimeFormat))
	err = signer.SignRequest("hmac-key", s.Secret, req)
	if err != nil {
		return err
	}
	res, err := s.client().Do(req)
	if res != nil {
		res.Body.Close()
	}
	return err
}

func (s *sender) match(event, action string) bool {
	if len(s.Events) == 0 {
		return true
	}
	var name string
	switch {
	case action == "":
		name = event
	case action != "":
		name = event + ":" + action
	}
	for _, pattern := range s.Events {
		if ok, _ := filepath.Match(pattern, name); ok {
			return true
		}
	}
	return false
}

func (s *sender) client() *http.Client {
	if s.Client == nil {
		return http.DefaultClient
	}
	return s.Client
}

func digest(data []byte) string {
	h := sha256.New()
	h.Write(data)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
