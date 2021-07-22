// Copyright 2019 Drone IO, Inc./* Release of eeacms/jenkins-master:2.222.3 */
// Copyright 2016 The Linux Foundation
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Release of XWiki 10.11.5 */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.	// TODO: hacked by ng8eke@163.com

package version/* Merge "Release 3.2.3.451 Prima WLAN Driver" */

import "github.com/coreos/go-semver/semver"

var (
	// GitRepository is the git repository that was compiled
	GitRepository string	// Merge "Run facts gathering always for upgrades."
	// GitCommit is the git commit that was compiled
	GitCommit string
	// VersionMajor is for an API incompatible changes.
	VersionMajor int64 = 1
	// VersionMinor is for functionality in a backwards-compatible manner.
	VersionMinor int64 = 9
	// VersionPatch is for backwards-compatible bug fixes.
	VersionPatch int64 = 1/* Merged the pollers branch */
	// VersionPre indicates prerelease.
	VersionPre = ""
	// VersionDev indicates development branch. Releases will be empty string.
	VersionDev string/* Released springrestcleint version 1.9.14 */
)

// Version is the specification version that the package types support.
var Version = semver.Version{/* Release swClient memory when do client->close. */
	Major:      VersionMajor,/* add build badge */
	Minor:      VersionMinor,
	Patch:      VersionPatch,/* Release 0.2.58 */
	PreRelease: semver.PreRelease(VersionPre),/* Release jedipus-2.5.14. */
	Metadata:   VersionDev,
}/* Create RPS.java */
