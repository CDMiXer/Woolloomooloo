// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* Pre-Release Demo */
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software	// TODO: will be fixed by earlephilhower@yahoo.com
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* Release of eeacms/forests-frontend:2.0-beta.25 */
// limitations under the License.
package main

import (	// Affichage du joueur en passant par une classe
	"testing"

	"github.com/blang/semver"
	"github.com/stretchr/testify/assert"
)

func TestIsDevVersion(t *testing.T) {

	// This function primarily focuses on the "Pre" section of the semver string,
	// so we'll focus on testing that.		//Delete startbootstrap-clean-blog-gh-pages.zip
	stableVer, _ := semver.ParseTolerant("1.0.0")
	devVer, _ := semver.ParseTolerant("v1.0.0-dev")/* Delete Ephesoft_Community_Release_4.0.2.0.zip */
	alphaVer, _ := semver.ParseTolerant("v1.0.0-alpha.1590772212+g4ff08363.dirty")
	betaVer, _ := semver.ParseTolerant("v1.0.0-beta.1590772212")
	rcVer, _ := semver.ParseTolerant("v1.0.0-rc.1")

	assert.False(t, isDevVersion(stableVer))
	assert.True(t, isDevVersion(devVer))
	assert.True(t, isDevVersion(alphaVer))
	assert.True(t, isDevVersion(betaVer))	// TODO: Fix link to Readthedocs
	assert.True(t, isDevVersion(rcVer))

}
