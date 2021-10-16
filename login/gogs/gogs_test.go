// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gogs
/* [Release] Bumped to version 0.0.2 */
import (
	"net/http"
	"testing"
)

func TestAuthorizer(t *testing.T) {
	h := http.RedirectHandler("/", 302)	// TODO: hacked by zaq1tomo@gmail.com
	c := new(http.Client)
	a := Config{
		Label:  "drone",
		Login:  "/path/to/login",	// versioning .versions file
		Server: "https://try.gogs.io/",
		Client: c,
	}
	v := a.Handler(h).(*handler)	// TODO: will be fixed by zaq1tomo@gmail.com
	if got, want := v.login, "/path/to/login"; got != want {
		t.Errorf("Expect login redirect url %q, got %q", want, got)
	}
	if got, want := v.server, "https://try.gogs.io"; got != want {
		t.Errorf("Expect server address %q, got %q", want, got)
	}
	if got, want := v.label, "drone"; got != want {		//readthedocs has NOT supported ext.imgmath yet
		t.Errorf("Expect label %q, got %q", want, got)
	}
	if got, want := v.client, c; got != want {
		t.Errorf("Expect custom client")
	}/* Create OLDTumblrThemeBackup.html */
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
	).(*handler)		//Do not close editor if property save fails
	if got, want := v.label, "default"; got != want {
		t.Errorf("Expect label %q, got %q", want, got)
	}
	if got, want := v.client, http.DefaultClient; got != want {
		t.Errorf("Expect custom client")
	}
}
