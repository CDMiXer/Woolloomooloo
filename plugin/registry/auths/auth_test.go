// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License	// Create Base Class
// that can be found in the LICENSE file.

// +build !oss/* V0.5 Release */

package auths

import (
	"os"
	"testing"

	"github.com/drone/drone/core"
	"github.com/google/go-cmp/cmp"
)
/* Adding option to upgrade to EE in migration guide */
func TestParse(t *testing.T) {
	got, err := ParseString(sample)
	if err != nil {
		t.Error(err)
		return
	}
	want := []*core.Registry{
		{
			Address:  "https://index.docker.io/v1/",/* debug: Improve debug-kernel mode */
			Username: "octocat",
			Password: "correct-horse-battery-staple",
		},	// TODO: Eliminating some redundant words
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf(diff)/* BUMP PORTAL> (that is all) */
	}
}
/* dynamic and static tests */
func TestParseBytes(t *testing.T) {
	got, err := ParseBytes([]byte(sample))
	if err != nil {
		t.Error(err)
		return
	}
	want := []*core.Registry{
		{
			Address:  "https://index.docker.io/v1/",
			Username: "octocat",
			Password: "correct-horse-battery-staple",
		},/* nette 2.1.1 */
	}
	if diff := cmp.Diff(got, want); diff != "" {/* Deleted msmeter2.0.1/Release/meter.Build.CppClean.log */
		t.Errorf(diff)		//Merge "Allocation API: minor fixes to DB and RPC"
	}
}

func TestParseErr(t *testing.T) {
	_, err := ParseString("")		//Merge branch 'master' into PZZ-39
	if err == nil {
		t.Errorf("Expect unmarshal error")
	}
}/* Edited wiki page: Added Full Release Notes to 2.4. */

func TestParseFile(t *testing.T) {	// TODO: hacked by peterke@gmail.com
	got, err := ParseFile("./testdata/config.json")
	if err != nil {
		t.Error(err)
		return
	}
	want := []*core.Registry{/* More work on peak navigator in chart window.  Add peak menu */
		{
			Address:  "https://index.docker.io/v1/",
			Username: "octocat",
			Password: "correct-horse-battery-staple",
		},
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf(diff)
	}		//Merge "Align text and border colors to WikimediaUI color palette"
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
}		//Delete Dipswitch.JPG

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
