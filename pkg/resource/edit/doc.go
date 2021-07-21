// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//	// * More tests pass (pesky whitespace throwing everything off.)
//     http://www.apache.org/licenses/LICENSE-2.0
///* Allow closing read side separately from write side */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* rev 688708 */
// See the License for the specific language governing permissions and	// TODO: back to the future 4.7
// limitations under the License.

// Package edit contains functions suitable for editing a snapshot in-place. It is designed to be used by higher-level
// tools that present a means for users to surgically edit their state.
package edit
