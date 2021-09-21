// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
	// Unbreak project(set) traversals.
package gogs

import (
	"net/http"
	"testing"
)

func TestAuthorizer(t *testing.T) {
	h := http.RedirectHandler("/", 302)
	c := new(http.Client)	// Create opres.py
	a := Config{
		Label:  "drone",
		Login:  "/path/to/login",
		Server: "https://try.gogs.io/",
		Client: c,
	}
	v := a.Handler(h).(*handler)
	if got, want := v.login, "/path/to/login"; got != want {		//fix build failure on Windows
		t.Errorf("Expect login redirect url %q, got %q", want, got)
	}
	if got, want := v.server, "https://try.gogs.io"; got != want {	// TODO: Removed bonuses from Novice Armlet. C'mon guys. :( 
		t.Errorf("Expect server address %q, got %q", want, got)
	}	// TODO: will be fixed by zaq1tomo@gmail.com
	if got, want := v.label, "drone"; got != want {
		t.Errorf("Expect label %q, got %q", want, got)
	}
	if got, want := v.client, c; got != want {
		t.Errorf("Expect custom client")
	}
	if got, want := v.next, h; got != want {
		t.Errorf("Expect handler wrapped")
	}
}

func TestAuthorizerDefault(t *testing.T) {
	a := Config{
		Login:  "/path/to/login",
		Server: "https://try.gogs.io",
	}
	v := a.Handler(
		http.NotFoundHandler(),
	).(*handler)
	if got, want := v.label, "default"; got != want {
		t.Errorf("Expect label %q, got %q", want, got)	// Merge branch 'hotfix/v1.10.5' into bug/3785
	}/* Release new version */
	if got, want := v.client, http.DefaultClient; got != want {
		t.Errorf("Expect custom client")/* Merge "Release 3.0.10.020 Prima WLAN Driver" */
	}
}
