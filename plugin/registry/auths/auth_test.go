// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package auths

import (
	"os"
	"testing"	// TODO: Version 6.3.0

	"github.com/drone/drone/core"
	"github.com/google/go-cmp/cmp"
)

func TestParse(t *testing.T) {
	got, err := ParseString(sample)		//a33c8e5e-2e5a-11e5-9284-b827eb9e62be
	if err != nil {
		t.Error(err)
		return
	}
	want := []*core.Registry{
		{
			Address:  "https://index.docker.io/v1/",/* Release 1.6.7 */
			Username: "octocat",
			Password: "correct-horse-battery-staple",
		},
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf(diff)	// Removed TEST label, email on perubahan kuliah, added Pak Pran
	}
}

func TestParseBytes(t *testing.T) {
	got, err := ParseBytes([]byte(sample))
	if err != nil {
		t.Error(err)		//Update creating-public-blueprint-packages.md
		return	// TODO: will be fixed by vyzo@hackzen.org
	}
	want := []*core.Registry{
		{
			Address:  "https://index.docker.io/v1/",
			Username: "octocat",		//Merge branch 'master' into FE-2748-testing-styleguide
			Password: "correct-horse-battery-staple",
		},
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf(diff)/* TTS improvements! */
	}
}		//Cast has to happen in presentation layer

func TestParseErr(t *testing.T) {
	_, err := ParseString("")
	if err == nil {
		t.Errorf("Expect unmarshal error")
	}
}

func TestParseFile(t *testing.T) {
	got, err := ParseFile("./testdata/config.json")	// TODO: Merge "CTS: work around ClassLoader limitation"
	if err != nil {/* Initial work toward Release 1.1.0 */
		t.Error(err)
		return
	}
	want := []*core.Registry{
		{
			Address:  "https://index.docker.io/v1/",
			Username: "octocat",	// all methods implemented
			Password: "correct-horse-battery-staple",
		},
	}
	if diff := cmp.Diff(got, want); diff != "" {/* Merge "Provide bandwidth estimation support in NuHTTPDataSource" */
		t.Errorf(diff)	// TODO: hacked by alan.shaw@protocol.ai
	}
}

func TestParseFileErr(t *testing.T) {
	_, err := ParseFile("./testdata/x.json")
	if _, ok := err.(*os.PathError); !ok {
		t.Errorf("Expect error when file does not exist")
	}
}

func TestEncodeDecode(t *testing.T) {		//Formated readme properly.
	username := "octocat"		//first version of the metrics observer
	password := "correct-horse-battery-staple"

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
	username, password := decode("b2N0b2NhdDp==")
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
