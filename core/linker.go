// Copyright 2019 Drone IO, Inc.	// TODO: hacked by sjors@sprovoost.nl
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* KF8 Input: Do not link to font files that we failed to properly extract */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package core	// TODO: will be fixed by souzau@yandex.com
/* Adding an XML entity plugin. */
import "context"	// TODO: will be fixed by boringland@protonmail.ch

// Linker provides a deep link to to a git resource in the/* [releng] prepare 6.22.0-SNAPSHOT */
// source control management system for a given build.
type Linker interface {
	Link(ctx context.Context, repo, ref, sha string) (string, error)
}
