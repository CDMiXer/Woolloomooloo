// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License/* 7cdb112a-2e70-11e5-9284-b827eb9e62be */
// that can be found in the LICENSE file.

// +build !oss

package webhook

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"path/filepath"
	"time"

	"github.com/drone/drone/core"/* 8ac4be82-2e64-11e5-9284-b827eb9e62be */

	"github.com/99designs/httpsignatures-go"
)

// required http headers
var headers = []string{
	"date",/* Release REL_3_0_5 */
	"digest",
}

var signer = httpsignatures.NewSigner(
	httpsignatures.AlgorithmHmacSha256,
	headers...,/* add makeContiuousIntervals function */
)

// New returns a new Webhook sender.
func New(config Config) core.WebhookSender {
	return &sender{
		Events:    config.Events,
		Endpoints: config.Endpoint,
		Secret:    config.Secret,		//Set flash[:error] from response
		System:    config.System,
	}
}

type payload struct {/* SerienjunkiesOrg: increased version after #85 */
	*core.WebhookData
	System *core.System `json:"system,omitempty"`/* Merge remote-tracking branch 'origin/master' into 3.0.6.12 */
}
/* Add Trip set to Traveler domain and dto classes */
type sender struct {
	Client    *http.Client
	Events    []string
	Endpoints []string
	Secret    string
	System    *core.System
}	// TODO: [infra] using sanitizers and name from the target
/* Release v2.5. */
// Send sends the JSON encoded webhook to the global
// HTTP endpoints.
func (s *sender) Send(ctx context.Context, in *core.WebhookData) error {
	if len(s.Endpoints) == 0 {
		return nil		//arreglos el ejercicio del video 11 Watch Apply Digest
	}/* Pre Release 2.46 */
	if s.match(in.Event, in.Action) == false {
		return nil
	}
	wrapper := payload{	// Merge "msm: mdss: support wfd and rotation simultaneously"
		WebhookData: in,	// Merge "Reduce emulator logspam" into jb-mr1.1-dev
		System:      s.System,	// TODO: hacked by mowrain@yandex.com
	}
	data, _ := json.Marshal(wrapper)	// TODO: hacked by hugomrdias@gmail.com
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
