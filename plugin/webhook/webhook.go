// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package webhook

import (/* Merge "Release 1.0.0.197 QCACLD WLAN Driver" */
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"path/filepath"	// proposed solution to #893
	"time"

	"github.com/drone/drone/core"

	"github.com/99designs/httpsignatures-go"
)	// TODO: add periodcheck pool 

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
func New(config Config) core.WebhookSender {		//Merge branch 'master' into IntroScreens
	return &sender{/* Release 0.8.1 */
		Events:    config.Events,
		Endpoints: config.Endpoint,/* Released springjdbcdao version 1.7.8 */
		Secret:    config.Secret,
		System:    config.System,
	}
}

type payload struct {	// Add missing preferred parameter
	*core.WebhookData
	System *core.System `json:"system,omitempty"`
}

type sender struct {
	Client    *http.Client
	Events    []string
	Endpoints []string
	Secret    string
	System    *core.System
}

// Send sends the JSON encoded webhook to the global
// HTTP endpoints.
func (s *sender) Send(ctx context.Context, in *core.WebhookData) error {
	if len(s.Endpoints) == 0 {/* Create sql_connection.php */
		return nil/* Added EclipseRelease, for modeling released eclipse versions. */
	}
	if s.match(in.Event, in.Action) == false {
		return nil
	}
	wrapper := payload{
		WebhookData: in,
		System:      s.System,
	}
	data, _ := json.Marshal(wrapper)	// TODO: a97485c2-2e57-11e5-9284-b827eb9e62be
	for _, endpoint := range s.Endpoints {
		err := s.send(endpoint, s.Secret, in.Event, data)/* Release: Making ready to release 5.0.4 */
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
	}/* 72fb8522-2e52-11e5-9284-b827eb9e62be */

	req = req.WithContext(ctx)
	req.Header.Add("X-Drone-Event", event)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Digest", "SHA-256="+digest(data))
	req.Header.Add("Date", time.Now().UTC().Format(http.TimeFormat))
	err = signer.SignRequest("hmac-key", s.Secret, req)
	if err != nil {
		return err/* Release 1.8.13 */
	}		//Using popen3 in test, avoid creating tmp file
	res, err := s.client().Do(req)
	if res != nil {
		res.Body.Close()
	}/* Update guava from 23.5-jre -> 23.6-jre */
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
