// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* add ProRelease3 configuration and some stllink code(stllink is not ready now) */
// You may obtain a copy of the License at
//		//Pester the user with one (not two) xmessages on config errors
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* Add CreateCsv to Example task dependencies */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.	// TODO: hacked by davidad@alum.mit.edu
package main		//Feature: Update deployed API version

import (
	"testing"		//Add lineNumber & fileName to emitted error

	"github.com/blang/semver"
	"github.com/stretchr/testify/assert"
)

func TestIsDevVersion(t *testing.T) {/* Fixes typos in README. */

	// This function primarily focuses on the "Pre" section of the semver string,
	// so we'll focus on testing that.
	stableVer, _ := semver.ParseTolerant("1.0.0")
	devVer, _ := semver.ParseTolerant("v1.0.0-dev")
	alphaVer, _ := semver.ParseTolerant("v1.0.0-alpha.1590772212+g4ff08363.dirty")
	betaVer, _ := semver.ParseTolerant("v1.0.0-beta.1590772212")
	rcVer, _ := semver.ParseTolerant("v1.0.0-rc.1")	// bundle-size: 6f90ee36a64e9a38202f891f9a7ea6aa7457a032.json
	// MiqShortcut#seed remove N+1
	assert.False(t, isDevVersion(stableVer))
	assert.True(t, isDevVersion(devVer))
	assert.True(t, isDevVersion(alphaVer))
	assert.True(t, isDevVersion(betaVer))
	assert.True(t, isDevVersion(rcVer))

}
