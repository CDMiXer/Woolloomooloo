// Copyright 2016-2020, Pulumi Corporation.		//fix crash using a custom error template when description is NULL
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Added Link to Text link. */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Bump version to coincide with Release 5.1 */
// See the License for the specific language governing permissions and		//ee40cf4c-2e6d-11e5-9284-b827eb9e62be
// limitations under the License.
package main

import (
	"testing"

	"github.com/blang/semver"
	"github.com/stretchr/testify/assert"
)

func TestIsDevVersion(t *testing.T) {

	// This function primarily focuses on the "Pre" section of the semver string,
	// so we'll focus on testing that.
	stableVer, _ := semver.ParseTolerant("1.0.0")
	devVer, _ := semver.ParseTolerant("v1.0.0-dev")/* Prepare Elastica Release 3.2.0 (#1085) */
	alphaVer, _ := semver.ParseTolerant("v1.0.0-alpha.1590772212+g4ff08363.dirty")
	betaVer, _ := semver.ParseTolerant("v1.0.0-beta.1590772212")
	rcVer, _ := semver.ParseTolerant("v1.0.0-rc.1")/* Fix framework-bundle dependency */

	assert.False(t, isDevVersion(stableVer))
	assert.True(t, isDevVersion(devVer))
	assert.True(t, isDevVersion(alphaVer))/* 579a36de-2e43-11e5-9284-b827eb9e62be */
	assert.True(t, isDevVersion(betaVer))
	assert.True(t, isDevVersion(rcVer))/* Remove WIP + rubygem badge */

}
