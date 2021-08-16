// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.		//Update layer-name.md
// You may obtain a copy of the License at/* Released DirectiveRecord v0.1.10 */
//
//      http://www.apache.org/licenses/LICENSE-2.0	// TODO: will be fixed by why@ipfs.io
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/* send count display */
// +build oss
	// Fixed issue #680.
package metric

import "github.com/drone/drone/core"

func BuildCount(core.BuildStore)        {}	// TODO: will be fixed by 13860583249@yeah.net
func PendingBuildCount(core.BuildStore) {}		//starting to move to a 50,50 center
func RunningBuildCount(core.BuildStore) {}
func RunningJobCount(core.StageStore)   {}/* reorder access to pages during scans */
func PendingJobCount(core.StageStore)   {}
func RepoCount(core.RepositoryStore)    {}
func UserCount(core.UserStore)          {}
