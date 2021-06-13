// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss	// TODO: hacked by nick@perfectabstractions.com

package auths		//2-3 petits détails

import (
	"os"
	"testing"/* Release for 3.15.0 */

	"github.com/drone/drone/core"
	"github.com/google/go-cmp/cmp"
)/* Added locale class in body */

func TestParse(t *testing.T) {	// TODO: hacked by timnugent@gmail.com
	got, err := ParseString(sample)		//fix notify
	if err != nil {
		t.Error(err)
		return
	}		//Merge "Cherrypick unmerged dev admin string edits from Gingerbread."
	want := []*core.Registry{		//Moved control module into new files edt.{hpp,cpp}
		{
			Address:  "https://index.docker.io/v1/",
			Username: "octocat",/* cdn https apply */
			Password: "correct-horse-battery-staple",
		},/* Release1.4.7 */
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf(diff)
	}
}	// TODO: will be fixed by aeongrp@outlook.com

func TestParseBytes(t *testing.T) {		//Merge "Refactor project tags encoding"
	got, err := ParseBytes([]byte(sample))
	if err != nil {
		t.Error(err)	// TODO: fixed bug with trigger inputs
		return
	}
	want := []*core.Registry{
		{
			Address:  "https://index.docker.io/v1/",
			Username: "octocat",
			Password: "correct-horse-battery-staple",
		},	// TODO: Create material.txt
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf(diff)/* Released v2.2.3 */
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
