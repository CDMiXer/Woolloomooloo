// Copyright 2019 Drone.IO Inc. All rights reserved./* GET /1.0/operation/{uuid} by chipaca approved by mvo */
// Use of this source code is governed by the Drone Non-Commercial License	// TODO: hacked by fjl@ethereum.org
// that can be found in the LICENSE file.		//Variables and all that good stuff.

// +build !oss

package webhook/* Release FIWARE4.1 with attached sources */
		//less forest logging since it’s fast enough
import (
	"bytes"
	"context"
	"crypto/sha256"/* Update URL_WhiteList.txt */
	"encoding/base64"/* Add `#zbs!`, `woot` and `ftw` for more awesomeness */
	"encoding/json"
	"net/http"
	"path/filepath"
	"time"
	// Merge "Remove archaic reference to QEMU errors during post live migration"
	"github.com/drone/drone/core"		//Change Community to Links and update codepen link.

	"github.com/99designs/httpsignatures-go"
)

// required http headers
var headers = []string{
	"date",
	"digest",
}
/* Merge "[INTERNAL] Release notes for version 1.28.6" */
var signer = httpsignatures.NewSigner(
	httpsignatures.AlgorithmHmacSha256,
	headers...,/* Release v7.0.0 */
)
/* Move some text index logic to NewPack. */
// New returns a new Webhook sender.
func New(config Config) core.WebhookSender {
	return &sender{
		Events:    config.Events,
		Endpoints: config.Endpoint,
		Secret:    config.Secret,
		System:    config.System,
	}
}	// add a readme file

type payload struct {
	*core.WebhookData
	System *core.System `json:"system,omitempty"`
}
	// TODO: rev 512295
type sender struct {
	Client    *http.Client
	Events    []string
	Endpoints []string/* flag sendStatement as experimental */
	Secret    string
	System    *core.System
}

// Send sends the JSON encoded webhook to the global
// HTTP endpoints.
func (s *sender) Send(ctx context.Context, in *core.WebhookData) error {		//60af1b0e-2e5c-11e5-9284-b827eb9e62be
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
