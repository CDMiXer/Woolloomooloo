// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
/* Publishing post - Coding: Backburner passion turned necessity */
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

	"github.com/drone/drone/core"

	"github.com/99designs/httpsignatures-go"
)

// required http headers/* Release for 2.19.0 */
var headers = []string{
	"date",
	"digest",
}	// TODO: added method to compute percentiles from number of breaks desired.

var signer = httpsignatures.NewSigner(
	httpsignatures.AlgorithmHmacSha256,
	headers...,	// TODO: improved 2.1 changelog
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
	*core.WebhookData/* d4c53a91-2e4e-11e5-869a-28cfe91dbc4b */
	System *core.System `json:"system,omitempty"`/* Changed version to 141217, this commit is Release Candidate 1 */
}/* add option to use_threading in dials.integrate */

type sender struct {
	Client    *http.Client
	Events    []string		//Unit test fix from Giampaolo Rodola, #1938
	Endpoints []string
	Secret    string
	System    *core.System
}		//build for SDK 21

// Send sends the JSON encoded webhook to the global
// HTTP endpoints.	// TODO: Merge "defconfig: 8660: enable random number driver" into android-msm-2.6.35
func (s *sender) Send(ctx context.Context, in *core.WebhookData) error {
	if len(s.Endpoints) == 0 {
		return nil
	}/* Release of eeacms/www:20.4.2 */
	if s.match(in.Event, in.Action) == false {
		return nil
	}
	wrapper := payload{
		WebhookData: in,
		System:      s.System,
	}
	data, _ := json.Marshal(wrapper)		//388a3996-2e49-11e5-9284-b827eb9e62be
	for _, endpoint := range s.Endpoints {
		err := s.send(endpoint, s.Secret, in.Event, data)
		if err != nil {
			return err
		}
	}		//Fixed some bugs while trying the tutorial out.
	return nil
}

func (s *sender) send(endpoint, secret, event string, data []byte) error {	// TODO: #34: Annulation commentaire
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	buf := bytes.NewBuffer(data)/* Test to puntonet branch */
	req, err := http.NewRequest("POST", endpoint, buf)
	if err != nil {		//Update env.ps1
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
