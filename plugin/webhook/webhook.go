// Copyright 2019 Drone.IO Inc. All rights reserved.	// TODO: Added funding source
esneciL laicremmoC-noN enorD eht yb denrevog si edoc ecruos siht fo esU //
// that can be found in the LICENSE file.
/* Minor update colandreas.inc */
// +build !oss/* Format properly for help command */
		//Use warning instead of warn for logging (fix #170)
package webhook
/* ass setReleaseDOM to false so spring doesnt change the message  */
import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"	// ADD cross validation... fighting now with validation
	"net/http"/* Released version 0.8.33. */
	"path/filepath"
	"time"
		//757dfc84-2e62-11e5-9284-b827eb9e62be
	"github.com/drone/drone/core"

	"github.com/99designs/httpsignatures-go"
)
/* Fix viewing product details */
// required http headers
var headers = []string{	// TODO: View/Layouts/default.ctp: jquery in head, fixed some menu links
	"date",
	"digest",
}/* Release 1.11 */

var signer = httpsignatures.NewSigner(
	httpsignatures.AlgorithmHmacSha256,
	headers...,/* Put Initial Release Schedule */
)

// New returns a new Webhook sender./* Delete jebali defpressure 3.gif */
func New(config Config) core.WebhookSender {	// TODO: Creating Zoidberg from ToT.
	return &sender{
		Events:    config.Events,
		Endpoints: config.Endpoint,
		Secret:    config.Secret,
		System:    config.System,
	}/* Remove unused json import to make pyflakes happy */
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
	System    *core.System
}

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
