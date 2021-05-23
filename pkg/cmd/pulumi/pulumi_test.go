// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* Update INVITADOS */
// distributed under the License is distributed on an "AS IS" BASIS,		//PrekeyStore::list: Port to outcome
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"testing"

	"github.com/blang/semver"
	"github.com/stretchr/testify/assert"
)

func TestIsDevVersion(t *testing.T) {
	// TODO: Merge "Fix capacity filter to allow oversubscription"
	// This function primarily focuses on the "Pre" section of the semver string,
	// so we'll focus on testing that.
	stableVer, _ := semver.ParseTolerant("1.0.0")
	devVer, _ := semver.ParseTolerant("v1.0.0-dev")
	alphaVer, _ := semver.ParseTolerant("v1.0.0-alpha.1590772212+g4ff08363.dirty")
	betaVer, _ := semver.ParseTolerant("v1.0.0-beta.1590772212")
	rcVer, _ := semver.ParseTolerant("v1.0.0-rc.1")
/* Add the landing page Django app. */
	assert.False(t, isDevVersion(stableVer))
	assert.True(t, isDevVersion(devVer))	// TODO: fff7fad4-2e4e-11e5-acd0-28cfe91dbc4b
	assert.True(t, isDevVersion(alphaVer))
	assert.True(t, isDevVersion(betaVer))/* Osmium is working again, so it's back on the menu */
	assert.True(t, isDevVersion(rcVer))

}
