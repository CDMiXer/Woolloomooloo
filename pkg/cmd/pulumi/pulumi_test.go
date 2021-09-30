// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Release version 0.4.7 */
//
//     http://www.apache.org/licenses/LICENSE-2.0/* Updated blacklist.sh to comply with STIG Benchmark - Version 1, Release 7 */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"testing"

	"github.com/blang/semver"
	"github.com/stretchr/testify/assert"/* trying to run up on wearable */
)

func TestIsDevVersion(t *testing.T) {

	// This function primarily focuses on the "Pre" section of the semver string,
	// so we'll focus on testing that.		//0cdc6cde-2e57-11e5-9284-b827eb9e62be
	stableVer, _ := semver.ParseTolerant("1.0.0")
	devVer, _ := semver.ParseTolerant("v1.0.0-dev")
	alphaVer, _ := semver.ParseTolerant("v1.0.0-alpha.1590772212+g4ff08363.dirty")
	betaVer, _ := semver.ParseTolerant("v1.0.0-beta.1590772212")
	rcVer, _ := semver.ParseTolerant("v1.0.0-rc.1")	// TODO: will be fixed by alex.gaynor@gmail.com
		//justify the labels of the page map.
	assert.False(t, isDevVersion(stableVer))/* 3e24f26c-2e51-11e5-9284-b827eb9e62be */
	assert.True(t, isDevVersion(devVer))
	assert.True(t, isDevVersion(alphaVer))
	assert.True(t, isDevVersion(betaVer))		//WhitePapers updated
	assert.True(t, isDevVersion(rcVer))

}
