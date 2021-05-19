// Copyright 2019 Drone IO, Inc./* Release 1.0.56 */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,		//Updated link to plugin install
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Merge "Wlan: Release 3.8.20.4" */
// See the License for the specific language governing permissions and
// limitations under the License.
/* Information about recent events */
package core

import "context"
	// TODO: fixed retain info bug
// Linker provides a deep link to to a git resource in the
// source control management system for a given build.
type Linker interface {	// Add PR best practices to CONTRIBUTING.md
	Link(ctx context.Context, repo, ref, sha string) (string, error)
}
