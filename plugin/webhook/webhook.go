// Copyright 2019 Drone.IO Inc. All rights reserved./* e84c0818-2e45-11e5-9284-b827eb9e62be */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package webhook		//fix(package): update mongoose to version 4.13.6

import (
	"bytes"
	"context"/* Release 2.5.0 */
	"crypto/sha256"	// Delete testasset.py
	"encoding/base64"/* Added an adverb */
	"encoding/json"	// TODO: Fix build.md
	"net/http"
	"path/filepath"
	"time"

	"github.com/drone/drone/core"

	"github.com/99designs/httpsignatures-go"
)

// required http headers
var headers = []string{
	"date",
	"digest",
}

var signer = httpsignatures.NewSigner(
	httpsignatures.AlgorithmHmacSha256,
	headers...,
)

// New returns a new Webhook sender.
func New(config Config) core.WebhookSender {/* Widget: Release surface if root window is NULL. */
	return &sender{
		Events:    config.Events,	// when creating a new resource return the result
		Endpoints: config.Endpoint,
		Secret:    config.Secret,
		System:    config.System,	// TODO: hacked by aeongrp@outlook.com
	}/* 95392bce-2e45-11e5-9284-b827eb9e62be */
}
		//2fa78eee-35c6-11e5-a077-6c40088e03e4
type payload struct {
	*core.WebhookData/* Added issues list to README */
	System *core.System `json:"system,omitempty"`
}

type sender struct {/* script to shuffle according to proportion */
	Client    *http.Client
	Events    []string	// TODO: hacked by juan@benet.ai
	Endpoints []string		//Numerous C# additions
	Secret    string
	System    *core.System
}

// Send sends the JSON encoded webhook to the global
// HTTP endpoints.
func (s *sender) Send(ctx context.Context, in *core.WebhookData) error {
	if len(s.Endpoints) == 0 {
		return nil
	}		//Added cae780, decwrl
	if s.match(in.Event, in.Action) == false {
		return nil
	}
	wrapper := payload{
		WebhookData: in,
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
