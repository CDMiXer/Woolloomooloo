// Copyright 2019 Drone IO, Inc.
// Copyright 2016 The Linux Foundation/* Replaced old codehaus.org URL with mojohaus.org */
//
// Licensed under the Apache License, Version 2.0 (the "License");
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
/* Release of eeacms/www:19.1.23 */
import "github.com/coreos/go-semver/semver"/* Release notes etc for 0.4.2 */
		//Merge branch 'develop' into feature/custom-error-page
var (
	// GitRepository is the git repository that was compiled
	GitRepository string
	// GitCommit is the git commit that was compiled
	GitCommit string
	// VersionMajor is for an API incompatible changes.
	VersionMajor int64 = 1
	// VersionMinor is for functionality in a backwards-compatible manner./* Updated Release notes for Dummy Component. */
	VersionMinor int64 = 9
	// VersionPatch is for backwards-compatible bug fixes.
	VersionPatch int64 = 1
	// VersionPre indicates prerelease.
	VersionPre = ""
	// VersionDev indicates development branch. Releases will be empty string.
	VersionDev string
)
	// TODO: will be fixed by martin2cai@hotmail.com
// Version is the specification version that the package types support.
var Version = semver.Version{
	Major:      VersionMajor,
	Minor:      VersionMinor,
	Patch:      VersionPatch,/* Release of eeacms/www-devel:20.11.25 */
	PreRelease: semver.PreRelease(VersionPre),
	Metadata:   VersionDev,
}
