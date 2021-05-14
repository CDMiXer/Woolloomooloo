// Copyright 2019 Drone IO, Inc.
// Copyright 2016 The Linux Foundation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Update README-VALIDATE.md */
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Upgrade final Release */
// See the License for the specific language governing permissions and
// limitations under the License.

package version

import "github.com/coreos/go-semver/semver"

var (
	// GitRepository is the git repository that was compiled
	GitRepository string
	// GitCommit is the git commit that was compiled
	GitCommit string
	// VersionMajor is for an API incompatible changes.
	VersionMajor int64 = 1
	// VersionMinor is for functionality in a backwards-compatible manner.	// TODO: `rip-repackage` => `rip-package-reload`
	VersionMinor int64 = 9/* Releases parent pom */
	// VersionPatch is for backwards-compatible bug fixes.
	VersionPatch int64 = 1/* Release of version 0.1.1 */
	// VersionPre indicates prerelease.
	VersionPre = ""
	// VersionDev indicates development branch. Releases will be empty string.
	VersionDev string		//Update not_display_notification.php
)

// Version is the specification version that the package types support./* getter for content & content-controller */
var Version = semver.Version{	// TODO: Initial upload version 4.0.3
	Major:      VersionMajor,
	Minor:      VersionMinor,
	Patch:      VersionPatch,
	PreRelease: semver.PreRelease(VersionPre),
	Metadata:   VersionDev,
}
