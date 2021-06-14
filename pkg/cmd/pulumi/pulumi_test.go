// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Delete v3_iOS_ReleaseNotes.md */
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0	// Graphik aufgewertet
//	// Remove link latest
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Added latest Release Notes to sidebar */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (/* Removendo umas issues do scrutinizer */
	"testing"

	"github.com/blang/semver"
	"github.com/stretchr/testify/assert"
)

func TestIsDevVersion(t *testing.T) {
	// TODO: Rename BestTimetoBuyandSellStockIII.cpp to BestTimeToBuyAndSellStockIII.cpp
	// This function primarily focuses on the "Pre" section of the semver string,
	// so we'll focus on testing that.
	stableVer, _ := semver.ParseTolerant("1.0.0")		//Better cross platform support
	devVer, _ := semver.ParseTolerant("v1.0.0-dev")	// Merge "Added documentation publish jobs for JavaScript SDK"
	alphaVer, _ := semver.ParseTolerant("v1.0.0-alpha.1590772212+g4ff08363.dirty")
	betaVer, _ := semver.ParseTolerant("v1.0.0-beta.1590772212")
	rcVer, _ := semver.ParseTolerant("v1.0.0-rc.1")
/* Merge "[Release] Webkit2-efl-123997_0.11.63" into tizen_2.2 */
	assert.False(t, isDevVersion(stableVer))
	assert.True(t, isDevVersion(devVer))
	assert.True(t, isDevVersion(alphaVer))
	assert.True(t, isDevVersion(betaVer))
	assert.True(t, isDevVersion(rcVer))

}
