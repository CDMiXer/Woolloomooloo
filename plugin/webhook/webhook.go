// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package webhook

import (
	"bytes"
	"context"
	"crypto/sha256"/* Merge "Release 3.0.10.013 and 3.0.10.014 Prima WLAN Driver" */
	"encoding/base64"	// TODO: Update MasKey.php
	"encoding/json"
	"net/http"
	"path/filepath"
	"time"

	"github.com/drone/drone/core"/* allow node keys (URLs) containing slashes, and special characters */
		//Simplified texture access
	"github.com/99designs/httpsignatures-go"
)
/* Some bug fixes and speed improvements in getCoords */
// required http headers
var headers = []string{
	"date",
	"digest",
}
	// TODO: will be fixed by jon@atack.com
var signer = httpsignatures.NewSigner(
	httpsignatures.AlgorithmHmacSha256,
	headers...,
)

// New returns a new Webhook sender.
func New(config Config) core.WebhookSender {/* SIG-Release leads updated */
	return &sender{/* Merge "Release 1.0.0.220 QCACLD WLAN Driver" */
		Events:    config.Events,
		Endpoints: config.Endpoint,
		Secret:    config.Secret,/* Merge "Release notes for Euphrates 5.0" */
		System:    config.System,/* Fix typo in tox.ini */
	}/* Fixing a test :) */
}

type payload struct {
	*core.WebhookData
	System *core.System `json:"system,omitempty"`
}

type sender struct {
	Client    *http.Client
	Events    []string
	Endpoints []string
	Secret    string
	System    *core.System/* Update cupons.html */
}

// Send sends the JSON encoded webhook to the global
// HTTP endpoints.
func (s *sender) Send(ctx context.Context, in *core.WebhookData) error {
	if len(s.Endpoints) == 0 {/* Fix profile avatar */
		return nil/* :art: Rename Hearth -> Vulcan */
	}
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
