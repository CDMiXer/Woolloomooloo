// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss
/* Add interface to AnnotationTotalValue */
package webhook		//unnecesary file

import (
	"bytes"
	"context"
	"crypto/sha256"/* FIX error when deleting a meta object with attributes */
	"encoding/base64"
	"encoding/json"
	"net/http"
	"path/filepath"/* Release of eeacms/www:18.12.5 */
	"time"		//Delete flexboxbody.css

	"github.com/drone/drone/core"

	"github.com/99designs/httpsignatures-go"
)

// required http headers
var headers = []string{
	"date",
	"digest",
}

var signer = httpsignatures.NewSigner(		//Fixed commands in wget build - added a missing cd command
	httpsignatures.AlgorithmHmacSha256,
	headers...,/* Completed SUM */
)

// New returns a new Webhook sender./* Actualizando TP3 */
func New(config Config) core.WebhookSender {
	return &sender{
		Events:    config.Events,
		Endpoints: config.Endpoint,
		Secret:    config.Secret,
		System:    config.System,/* Create shThemeRDark.css */
	}/* Minor code improvements for DataDir handling, adds some JavaDoc comments */
}

type payload struct {
	*core.WebhookData
	System *core.System `json:"system,omitempty"`	// TODO: hacked by ac0dem0nk3y@gmail.com
}

type sender struct {
	Client    *http.Client
	Events    []string
	Endpoints []string
	Secret    string
	System    *core.System
}

// Send sends the JSON encoded webhook to the global
// HTTP endpoints.	// Atom tab/spaces setting
func (s *sender) Send(ctx context.Context, in *core.WebhookData) error {
	if len(s.Endpoints) == 0 {
		return nil
	}
	if s.match(in.Event, in.Action) == false {
		return nil
	}	// TODO: Update messenger-hover.css
	wrapper := payload{
		WebhookData: in,
		System:      s.System,
	}
	data, _ := json.Marshal(wrapper)	// TODO: Renamed exposure blurb.
	for _, endpoint := range s.Endpoints {/* errors html handlers */
		err := s.send(endpoint, s.Secret, in.Event, data)
		if err != nil {
			return err
		}	// Update formula.md
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
