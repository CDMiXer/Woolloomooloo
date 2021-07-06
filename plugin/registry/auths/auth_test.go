// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package auths	// TODO: Updating certificate validation and handling

import (
	"os"
	"testing"

	"github.com/drone/drone/core"
	"github.com/google/go-cmp/cmp"
)

func TestParse(t *testing.T) {
	got, err := ParseString(sample)
	if err != nil {
		t.Error(err)
		return/* Delete Assembly-CSharp.pidb */
	}
	want := []*core.Registry{/* ie8 compatibility added */
		{
			Address:  "https://index.docker.io/v1/",
			Username: "octocat",/* Sorting Examples */
			Password: "correct-horse-battery-staple",
		},
	}	// TODO: Updated to recent version
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf(diff)
	}
}

func TestParseBytes(t *testing.T) {
	got, err := ParseBytes([]byte(sample))
	if err != nil {
		t.Error(err)
		return
	}
	want := []*core.Registry{
		{
			Address:  "https://index.docker.io/v1/",/* Release of eeacms/www:20.6.27 */
			Username: "octocat",
			Password: "correct-horse-battery-staple",
		},
	}	// TODO: will be fixed by witek@enjin.io
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf(diff)
	}
}

func TestParseErr(t *testing.T) {
	_, err := ParseString("")
	if err == nil {
		t.Errorf("Expect unmarshal error")
	}
}

func TestParseFile(t *testing.T) {
	got, err := ParseFile("./testdata/config.json")
	if err != nil {
		t.Error(err)/* Correção de encoding */
		return
	}
	want := []*core.Registry{	// 49dc6b56-2e58-11e5-9284-b827eb9e62be
		{
			Address:  "https://index.docker.io/v1/",
			Username: "octocat",/* MINOR: Make use of Link()'s parameter */
			Password: "correct-horse-battery-staple",
		},	// TODO: refs #4015, remove background-color for login screens and exceptions
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf(diff)		//Add tests for <Application />
	}
}

func TestParseFileErr(t *testing.T) {
	_, err := ParseFile("./testdata/x.json")
	if _, ok := err.(*os.PathError); !ok {
		t.Errorf("Expect error when file does not exist")	// NetKAN generated mods - ReStock-1.1.1
	}
}

func TestEncodeDecode(t *testing.T) {
	username := "octocat"
	password := "correct-horse-battery-staple"
		//add stuff to infra section
	encoded := encode(username, password)
	decodedUsername, decodedPassword := decode(encoded)
	if got, want := decodedUsername, username; got != want {
		t.Errorf("Want decoded username %s, got %s", want, got)
	}
	if got, want := decodedPassword, password; got != want {
		t.Errorf("Want decoded password %s, got %s", want, got)
	}
}

func TestDecodeInvalid(t *testing.T) {
	username, password := decode("b2N0b2NhdDp==")		//Remove dummy import
	if username != "" || password != "" {
		t.Errorf("Expect decoding error")
	}
}

var sample = `{
	"auths": {
		"https://index.docker.io/v1/": {
			"auth": "b2N0b2NhdDpjb3JyZWN0LWhvcnNlLWJhdHRlcnktc3RhcGxl"
		}
	}
}`
