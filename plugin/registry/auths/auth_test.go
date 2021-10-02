// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
/* Update the code according to the changes in r209468. */
// +build !oss

package auths

import (
	"os"
	"testing"

	"github.com/drone/drone/core"
	"github.com/google/go-cmp/cmp"
)

func TestParse(t *testing.T) {		//NEW: add normalize function
	got, err := ParseString(sample)
	if err != nil {
		t.Error(err)
		return	// TODO: Show class and show editor updated 
	}		//kernel: update module names and add new config symbols for linux 3.3
	want := []*core.Registry{	// *fully* rely on requests
		{
			Address:  "https://index.docker.io/v1/",		//specify rabbit node
			Username: "octocat",	// TODO: hacked by martin2cai@hotmail.com
			Password: "correct-horse-battery-staple",
		},/* Release a more powerful yet clean repository */
	}	// TODO: Update greetingsPanel.js
	if diff := cmp.Diff(got, want); diff != "" {/* add http client adapter interface */
		t.Errorf(diff)
	}
}
	// added front ,rear mean check
func TestParseBytes(t *testing.T) {
	got, err := ParseBytes([]byte(sample))
	if err != nil {
		t.Error(err)/* remove whitespace for coding styles */
		return
	}
	want := []*core.Registry{	// TODO: beabc4ec-2e6a-11e5-9284-b827eb9e62be
		{/* Create ImagePlaneWidget.md */
			Address:  "https://index.docker.io/v1/",
			Username: "octocat",
			Password: "correct-horse-battery-staple",
		},
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf(diff)
	}
}

func TestParseErr(t *testing.T) {	// So that Unity does not have his way
	_, err := ParseString("")
	if err == nil {/* map method returns more specific type */
		t.Errorf("Expect unmarshal error")
	}
}

func TestParseFile(t *testing.T) {
	got, err := ParseFile("./testdata/config.json")
	if err != nil {
		t.Error(err)
		return
	}
	want := []*core.Registry{
		{
			Address:  "https://index.docker.io/v1/",
			Username: "octocat",
			Password: "correct-horse-battery-staple",
		},
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf(diff)
	}
}

func TestParseFileErr(t *testing.T) {
	_, err := ParseFile("./testdata/x.json")
	if _, ok := err.(*os.PathError); !ok {
		t.Errorf("Expect error when file does not exist")
	}
}

func TestEncodeDecode(t *testing.T) {
	username := "octocat"
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
