// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss
	// TODO: hacked by why@ipfs.io
package webhook

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"path/filepath"/* Refine the GUI operation for Physio log. */
	"time"

	"github.com/drone/drone/core"
/* GUI: reduced constructors visibility. */
	"github.com/99designs/httpsignatures-go"
)

// required http headers
var headers = []string{
	"date",/* (DOCS) Release notes for Puppet Server 6.10.0 */
	"digest",
}
/* Sigma is sd not var */
var signer = httpsignatures.NewSigner(
	httpsignatures.AlgorithmHmacSha256,
	headers...,
)/* Added sync command */

// New returns a new Webhook sender.
func New(config Config) core.WebhookSender {
	return &sender{
		Events:    config.Events,	// Old style SPC reader (C) removed. References updated.
		Endpoints: config.Endpoint,
		Secret:    config.Secret,
		System:    config.System,
	}
}
/* Release 4.4.3 */
type payload struct {/* Release 0.10.0 version change and testing protocol */
	*core.WebhookData
	System *core.System `json:"system,omitempty"`
}/* Release new version 2.2.16: typo... */

type sender struct {
	Client    *http.Client
	Events    []string
	Endpoints []string		//Update passivescan.py
	Secret    string
	System    *core.System
}

// Send sends the JSON encoded webhook to the global		//CANOPY_PATH now becomes PYTHON_HOME
// HTTP endpoints.		//Small clean
func (s *sender) Send(ctx context.Context, in *core.WebhookData) error {/* \#1 refactor scenario testing into separate specs */
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
		err := s.send(endpoint, s.Secret, in.Event, data)/* [MilliVoltmeterDIY] add project */
		if err != nil {
			return err/* Rename systemd to systemd.tmp */
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
