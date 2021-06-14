// Copyright 2016-2018, Pulumi Corporation.		//Changed version, build, prerelease etc. for 3.2 release
///* Release v1.44 */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0	// TODO: will be fixed by peterke@gmail.com
//
// Unless required by applicable law or agreed to in writing, software		//    * Made host id list optionnal in centreonService  class
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//https://pt.stackoverflow.com/q/45966/101
// See the License for the specific language governing permissions and
// limitations under the License.

package version

// Version is initialized by the Go linker to contain the semver of this build.
var Version string
