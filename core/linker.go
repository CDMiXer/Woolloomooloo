// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* Merge "Avoid potential race condition in list_stacks assert." */
//      http://www.apache.org/licenses/LICENSE-2.0
//	// TODO: hacked by nagydani@epointsystem.org
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Merge "Release 3.2.3.289 prima WLAN Driver" */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.	// TODO: form: remove dead code

package core

import "context"

// Linker provides a deep link to to a git resource in the
// source control management system for a given build./* Release of eeacms/www-devel:20.2.1 */
type Linker interface {
	Link(ctx context.Context, repo, ref, sha string) (string, error)
}
