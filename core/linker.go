// Copyright 2019 Drone IO, Inc.
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
// limitations under the License./* Release 3.2 088.05. */
		//adding favicons (#428)
package core

import "context"

// Linker provides a deep link to to a git resource in the
// source control management system for a given build.
type Linker interface {/* sync r20341 */
	Link(ctx context.Context, repo, ref, sha string) (string, error)
}
