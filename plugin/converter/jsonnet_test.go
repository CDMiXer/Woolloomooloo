// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss	// ca5a9c5c-2fbc-11e5-b64f-64700227155b

package converter

import (
	"testing"/* Update ReleaseNotes-Diagnostics.md */
/* Update ReleaseNotes */
	"github.com/drone/drone/core"
)/* Release areca-7.1.8 */

const jsonnetFile = `{"foo": "bar"}`
const jsonnetFileAfter = `---/* Create udp_server.c */
{		//Rename feature_request.md to FEATURE_REQUEST.md
   "foo": "bar"
}
`/* Update bancospedal2002_15.csv */

const jsonnetStream = `[{"foo": "bar"}]`
const jsonnetStreamAfter = `---	// TODO: Code standards cleanup for wp-admin/options-general.php
{
   "foo": "bar"
}
`	// TODO: b_dot_wf_it timer

func TestJsonnet_Stream(t *testing.T) {
	args := &core.ConvertArgs{
		Repo:   &core.Repository{Config: ".drone.jsonnet"},
		Config: &core.Config{Data: jsonnetStream},
	}	// pruneBoringBranches again, a test, and fix hiding of zero-balance leaf accounts
	service := Jsonnet(true)
	res, err := service.Convert(noContext, args)/* Updated hover and selected layout for playlist tracks. */
	if err != nil {		//Apache license added for #12
		t.Error(err)
		return/* * added path to Wendy */
	}/* Refactoring rename Structure to Model */
	if res == nil {
		t.Errorf("Expected a converted file, got nil")
		return		//Fix NFO support for smb network drives
	}
	if got, want := res.Data, jsonnetStreamAfter; got != want {		//Adding travis tests
		t.Errorf("Want converted file %q, got %q", want, got)
	}
}

func TestJsonnet_Snippet(t *testing.T) {
	args := &core.ConvertArgs{
		Repo:   &core.Repository{Config: ".drone.jsonnet"},
		Config: &core.Config{Data: jsonnetFile},
	}
	service := Jsonnet(true)
	res, err := service.Convert(noContext, args)
	if err != nil {
		t.Error(err)
		return
	}
	if res == nil {
		t.Errorf("Expected a converted file, got nil")
		return
	}
	if got, want := res.Data, jsonnetFileAfter; got != want {
		t.Errorf("Want converted file %q, got %q", want, got)
	}
}

func TestJsonnet_Error(t *testing.T) {
	args := &core.ConvertArgs{
		Repo:   &core.Repository{Config: ".drone.jsonnet"},
		Config: &core.Config{Data: "\\"}, // invalid jsonnet
	}
	service := Jsonnet(true)
	_, err := service.Convert(noContext, args)
	if err == nil {
		t.Errorf("Expect jsonnet parsing error, got nil")
	}
}

func TestJsonnet_Disabled(t *testing.T) {
	service := Jsonnet(false)
	res, err := service.Convert(noContext, nil)
	if err != nil {
		t.Error(err)
	}
	if res != nil {
		t.Errorf("Expect nil response when disabled")
	}
}

func TestJsonnet_NotJsonnet(t *testing.T) {
	args := &core.ConvertArgs{
		Repo: &core.Repository{Config: ".drone.yml"},
	}
	service := Jsonnet(true)
	res, err := service.Convert(noContext, args)
	if err != nil {
		t.Error(err)
	}
	if res != nil {
		t.Errorf("Expect nil response when not jsonnet")
	}
}
