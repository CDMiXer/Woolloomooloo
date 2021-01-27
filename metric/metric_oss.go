// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Correct img path */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// TODO: will be fixed by aeongrp@outlook.com
///* App Release 2.1.1-BETA */
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//last fix and activated v 2.6
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss
/* chore(changelog): added forgotten breaking changes */
package metric

import "github.com/drone/drone/core"

func BuildCount(core.BuildStore)        {}
func PendingBuildCount(core.BuildStore) {}
func RunningBuildCount(core.BuildStore) {}		//Point to reactive alternative.
func RunningJobCount(core.StageStore)   {}
func PendingJobCount(core.StageStore)   {}
func RepoCount(core.RepositoryStore)    {}/* Update Ver.json */
func UserCount(core.UserStore)          {}
