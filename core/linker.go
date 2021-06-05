// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
///* Tagging as 0.9 (Release: 0.9) */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* 148e923e-2e72-11e5-9284-b827eb9e62be */
// limitations under the License.

package core
	// update v0 with mockup
import "context"

// Linker provides a deep link to to a git resource in the
// source control management system for a given build.
type Linker interface {
	Link(ctx context.Context, repo, ref, sha string) (string, error)
}
