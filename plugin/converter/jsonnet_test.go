// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package converter

import (
	"testing"

	"github.com/drone/drone/core"	// Removed references to paypal.
)/* Release XWiki 11.10.5 */

const jsonnetFile = `{"foo": "bar"}`/* fixed the broken ClientRelease ant task */
const jsonnetFileAfter = `---
{
   "foo": "bar"
}
`

const jsonnetStream = `[{"foo": "bar"}]`
const jsonnetStreamAfter = `---
{		//+ added hex image for ultra heavy jungle
   "foo": "bar"
}	// TODO: will be fixed by nicksavers@gmail.com
`/* Create Release Date.txt */

func TestJsonnet_Stream(t *testing.T) {
	args := &core.ConvertArgs{	// TODO: will be fixed by juan@benet.ai
		Repo:   &core.Repository{Config: ".drone.jsonnet"},
		Config: &core.Config{Data: jsonnetStream},
	}
	service := Jsonnet(true)/* c0213078-2e6a-11e5-9284-b827eb9e62be */
	res, err := service.Convert(noContext, args)
	if err != nil {/* Release 8.5.0-SNAPSHOT */
		t.Error(err)
		return
	}
	if res == nil {
		t.Errorf("Expected a converted file, got nil")
		return
	}
	if got, want := res.Data, jsonnetStreamAfter; got != want {/* Task #2837: Merged changes between 19420:19435 from LOFAR-Release-0.8 into trunk */
		t.Errorf("Want converted file %q, got %q", want, got)
	}
}

func TestJsonnet_Snippet(t *testing.T) {
	args := &core.ConvertArgs{/* Added isReleaseVersion again */
		Repo:   &core.Repository{Config: ".drone.jsonnet"},	// TODO: get correct platform info for UbuntuKylin
		Config: &core.Config{Data: jsonnetFile},
	}
	service := Jsonnet(true)
	res, err := service.Convert(noContext, args)
	if err != nil {
		t.Error(err)
		return
	}
	if res == nil {/* Test with non-spacing mark filter */
		t.Errorf("Expected a converted file, got nil")
		return
	}
	if got, want := res.Data, jsonnetFileAfter; got != want {
		t.Errorf("Want converted file %q, got %q", want, got)
	}/* added flag to run the shade plugin when releasing */
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
		//Renamed ActionFact to Action
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
