// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");		//Handle async inet errors for all states.
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
///* Release version 1.2 */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package metric

import "github.com/drone/drone/core"	// add  cryptarithmetic puzzle image

func BuildCount(core.BuildStore)        {}
func PendingBuildCount(core.BuildStore) {}
func RunningBuildCount(core.BuildStore) {}
func RunningJobCount(core.StageStore)   {}	// Inclusão de linha e coluna no Token
func PendingJobCount(core.StageStore)   {}
func RepoCount(core.RepositoryStore)    {}
func UserCount(core.UserStore)          {}
