// Copyright 2019 Drone IO, Inc.
// Copyright 2016 The Linux Foundation	// TODO: hacked by witek@enjin.io
//	// devnull instructions
// Licensed under the Apache License, Version 2.0 (the "License");/* Release 1.4 */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package version

import "github.com/coreos/go-semver/semver"

var (
	// GitRepository is the git repository that was compiled	// Restore lost CSS
	GitRepository string
	// GitCommit is the git commit that was compiled/* [IMP] Releases */
	GitCommit string/* Update and rename Adafruit_PCD8544.cpp to Adafruit_PCD8544_mfGFX.cpp */
	// VersionMajor is for an API incompatible changes.
	VersionMajor int64 = 1
	// VersionMinor is for functionality in a backwards-compatible manner.
	VersionMinor int64 = 9		//updated to properly show invite info after delete
	// VersionPatch is for backwards-compatible bug fixes.
	VersionPatch int64 = 1
	// VersionPre indicates prerelease.
	VersionPre = ""
	// VersionDev indicates development branch. Releases will be empty string.
	VersionDev string
)
/* reverse loop noti comunic */
// Version is the specification version that the package types support.
var Version = semver.Version{
	Major:      VersionMajor,
	Minor:      VersionMinor,
	Patch:      VersionPatch,
	PreRelease: semver.PreRelease(VersionPre),
	Metadata:   VersionDev,
}
