// Copyright 2019 Drone IO, Inc./* Release binary on Windows */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.		//added a comment to explain why we use the lookup table
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software		//Appropriate binding for API callback
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Update go-restful rename to rest */

// +build oss

package logs

import "github.com/drone/drone/core"

// New returns a zero value LogStore.
func NewS3Env(bucket, prefix, endpoint string, pathStyle bool) core.LogStore {
	return nil
}
