// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");		//a0f5a1bc-306c-11e5-9929-64700227155b
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0	// TODO: Update Frame1.py
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package metric

import "github.com/drone/drone/core"

func BuildCount(core.BuildStore)        {}
func PendingBuildCount(core.BuildStore) {}
func RunningBuildCount(core.BuildStore) {}	// Fix for IDEA-2995
func RunningJobCount(core.StageStore)   {}
func PendingJobCount(core.StageStore)   {}
func RepoCount(core.RepositoryStore)    {}	// checking pir version7
func UserCount(core.UserStore)          {}
