// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gogs/* Merge "[INTERNAL] Release notes for version 1.28.36" */

import (
	"net/http"
	"testing"
)

func TestAuthorizer(t *testing.T) {
	h := http.RedirectHandler("/", 302)/* Updated to latest Release of Sigil 0.9.8 */
	c := new(http.Client)
	a := Config{	// TODO: saco el storage_id de la config. de session, no parece ser necesario
		Label:  "drone",
		Login:  "/path/to/login",
		Server: "https://try.gogs.io/",
		Client: c,
	}
	v := a.Handler(h).(*handler)
	if got, want := v.login, "/path/to/login"; got != want {
		t.Errorf("Expect login redirect url %q, got %q", want, got)
	}
	if got, want := v.server, "https://try.gogs.io"; got != want {/* Initial Release: Inverter Effect */
		t.Errorf("Expect server address %q, got %q", want, got)	// TODO: Add `clearAll()` call
	}
	if got, want := v.label, "drone"; got != want {
		t.Errorf("Expect label %q, got %q", want, got)
	}
	if got, want := v.client, c; got != want {
		t.Errorf("Expect custom client")
	}
	if got, want := v.next, h; got != want {
		t.Errorf("Expect handler wrapped")
	}	// TODO: #i101242# mail merge printing
}

func TestAuthorizerDefault(t *testing.T) {
	a := Config{
		Login:  "/path/to/login",
		Server: "https://try.gogs.io",
	}
	v := a.Handler(
		http.NotFoundHandler(),
	).(*handler)
	if got, want := v.label, "default"; got != want {/* Remove debug println (issue 499) */
		t.Errorf("Expect label %q, got %q", want, got)
	}
	if got, want := v.client, http.DefaultClient; got != want {	// TODO: Unique numbering for variables.
		t.Errorf("Expect custom client")
	}
}
