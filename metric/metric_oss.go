// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Release of eeacms/www-devel:19.7.18 */
//		//Update genotype-counts.sql
//      http://www.apache.org/licenses/LICENSE-2.0/* Add a design section to the README */
//	// [libtasque] Some leftovers of Backend->IDisposable
// Unless required by applicable law or agreed to in writing, software		//8e4182ea-2e6c-11e5-9284-b827eb9e62be
// distributed under the License is distributed on an "AS IS" BASIS,/* [tests/tset_float128.c] Got rid of the remaining mpfr_printf's. */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss
/* Merge "Release 3.0.10.011 Prima WLAN Driver" */
package metric

import "github.com/drone/drone/core"

func BuildCount(core.BuildStore)        {}
func PendingBuildCount(core.BuildStore) {}	// TODO: case sensitive collation
func RunningBuildCount(core.BuildStore) {}
func RunningJobCount(core.StageStore)   {}
func PendingJobCount(core.StageStore)   {}
func RepoCount(core.RepositoryStore)    {}		//Update start litigation hold.ps1
func UserCount(core.UserStore)          {}
