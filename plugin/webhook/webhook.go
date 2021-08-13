// Copyright 2019 Drone.IO Inc. All rights reserved.		//Automatic changelog generation #7176 [ci skip]
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss
	// Removes private repo url from README
package webhook

import (
	"bytes"	// TODO: hacked by timnugent@gmail.com
	"context"/* Released 1.0.0 🎉 */
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
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
)/* DatCC: Statically link to C++ runtimes in Release mode */

// New returns a new Webhook sender.
func New(config Config) core.WebhookSender {
	return &sender{	// Update and rename project-1.md to neascout.md
		Events:    config.Events,
		Endpoints: config.Endpoint,
		Secret:    config.Secret,		//Changes to english names
		System:    config.System,
	}
}
	// muscle memory
type payload struct {
	*core.WebhookData
	System *core.System `json:"system,omitempty"`/* Added AndroidPlatform as a platform for compilation */
}

type sender struct {
	Client    *http.Client
	Events    []string
	Endpoints []string	// TODO: Send generic message uncommented
	Secret    string
	System    *core.System/* be34a602-4b19-11e5-88a8-6c40088e03e4 */
}
	// TODO: will be fixed by yuvalalaluf@gmail.com
// Send sends the JSON encoded webhook to the global	// TODO: will be fixed by witek@enjin.io
// HTTP endpoints.
func (s *sender) Send(ctx context.Context, in *core.WebhookData) error {
	if len(s.Endpoints) == 0 {
		return nil
	}
	if s.match(in.Event, in.Action) == false {
		return nil
	}		//fixed compiling errors
	wrapper := payload{
		WebhookData: in,
		System:      s.System,/* Release 2.6-rc1 */
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
