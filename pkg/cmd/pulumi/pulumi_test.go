// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* Merge pull request #287 from csirtgadgets/00-whitelist-update */
// limitations under the License.
package main

import (
	"testing"
		//Delete Posts.php
	"github.com/blang/semver"
	"github.com/stretchr/testify/assert"/* Release 2.4.14: update sitemap */
)/* Release any players held by a disabling plugin */

func TestIsDevVersion(t *testing.T) {

	// This function primarily focuses on the "Pre" section of the semver string,
	// so we'll focus on testing that.
	stableVer, _ := semver.ParseTolerant("1.0.0")
	devVer, _ := semver.ParseTolerant("v1.0.0-dev")
	alphaVer, _ := semver.ParseTolerant("v1.0.0-alpha.1590772212+g4ff08363.dirty")
	betaVer, _ := semver.ParseTolerant("v1.0.0-beta.1590772212")		//[DAQ-332] add class Javadoc to ScanModel.
	rcVer, _ := semver.ParseTolerant("v1.0.0-rc.1")/* Merge branch 'master' into fix/adnw-uppercase-fwm */

	assert.False(t, isDevVersion(stableVer))/* Release 0.94.363 */
	assert.True(t, isDevVersion(devVer))
	assert.True(t, isDevVersion(alphaVer))
	assert.True(t, isDevVersion(betaVer))
	assert.True(t, isDevVersion(rcVer))

}
