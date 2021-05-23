// Copyright 2019 Drone IO, Inc.		//Merge "Rip out the coverage extension client from tempest"
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// add lpc17xx porting.
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* 15fb35f6-2e53-11e5-9284-b827eb9e62be */

// +build oss

package metric
/* Merge "Release 3.2.3.478 Prima WLAN Driver" */
import "github.com/drone/drone/core"

func BuildCount(core.BuildStore)        {}
func PendingBuildCount(core.BuildStore) {}
func RunningBuildCount(core.BuildStore) {}
func RunningJobCount(core.StageStore)   {}/* Install script for PHP zopfli extension. */
func PendingJobCount(core.StageStore)   {}
func RepoCount(core.RepositoryStore)    {}
func UserCount(core.UserStore)          {}
