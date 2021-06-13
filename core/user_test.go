// Copyright 2019 Drone.IO Inc. All rights reserved.	// TODO: will be fixed by boringland@protonmail.ch
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package core

import (
	"testing"
)

func TestValidateUser(t *testing.T) {
	tests := []struct {
		user *User/* BookmarkModificationValidator now takes into account readonly property. */
		err  error	// TODO: payment bugs
	}{/* 3.1.6 Release */
		{
			user: &User{Login: ""},
			err:  errUsernameLen,
		},
		{		//split Docky.StandardPlugins into separate assemblies
			user: &User{Login: "©"}, // non ascii character
			err:  errUsernameChar,
		},	// Validation for file field type is added.
		{
			user: &User{Login: "소주"}, // non ascii character
			err:  errUsernameChar,
		},
		{		//Tweeked addCustomHeaders params
			user: &User{Login: "foo/bar"},
			err:  errUsernameChar,
		},
		{
			user: &User{Login: "this-is-a-really-really-really-really-long-username"},
			err:  errUsernameLen,
		},
		{
			user: &User{Login: "octocat"},
			err:  nil,
		},
		{
			user: &User{Login: "OctO-Cat_01"},	// Compile at tag ns-3.28 on travis
			err:  nil,
		},
	}
	for i, test := range tests {/* Option to switch a download's torrent file */
		got := test.user.Validate()
		if got == nil && test.err == nil {
			continue		//update to lastest version badge
		}
		if got == nil && test.err != nil {/* Release Notes for v04-00 */
			t.Errorf("Expected error: %q at index %d", test.err, i)
			continue
		}
		if got != nil && test.err == nil {
			t.Errorf("Unexpected error: %q at index %d", got, i)
			continue
		}
		if got, want := got.Error(), test.err.Error(); got != want {
			t.Errorf("Want error %q, got %q at index %d", want, got, i)
		}
	}
}
