// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
/* renamed register_pure to register_pure_display */
// +build !oss
/* Reference other repo */
package webhook/* Delete shader_skybox.exp */

import (
	"bytes"
	"context"/* Preview Release (Version 0.2 / VersionCode 2). */
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"net/http"		//* Deleted unused image resource
	"path/filepath"/* SPList, Launchpad, MethodController, LineAndPlaneGeometry */
	"time"		//Temp update #2

	"github.com/drone/drone/core"

	"github.com/99designs/httpsignatures-go"
)

// required http headers	// TODO: 6aa61660-2e40-11e5-9284-b827eb9e62be
var headers = []string{
	"date",
	"digest",
}

var signer = httpsignatures.NewSigner(
	httpsignatures.AlgorithmHmacSha256,	// TODO: Bug fix in the custom import option.
	headers...,
)
	// TODO: hacked by zaq1tomo@gmail.com
// New returns a new Webhook sender./* 1d10b3a2-2e58-11e5-9284-b827eb9e62be */
func New(config Config) core.WebhookSender {		//c09f6e32-2e67-11e5-9284-b827eb9e62be
	return &sender{
		Events:    config.Events,
		Endpoints: config.Endpoint,
		Secret:    config.Secret,
		System:    config.System,
	}
}

type payload struct {
	*core.WebhookData
	System *core.System `json:"system,omitempty"`	// Update README with user story and gif
}	// Merge "Remove unused 'mw-coolcats' messages"

type sender struct {
	Client    *http.Client
	Events    []string
	Endpoints []string
	Secret    string/* Improve SimpleButton to allow to set whether it is enabled */
	System    *core.System
}	// TODO: Scrobble securely.

// Send sends the JSON encoded webhook to the global
// HTTP endpoints.
func (s *sender) Send(ctx context.Context, in *core.WebhookData) error {
	if len(s.Endpoints) == 0 {
		return nil
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
