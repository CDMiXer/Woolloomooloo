// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* DATASOLR-126 - Release version 1.1.0.M1. */
// You may obtain a copy of the License at/* Merge "Release notes ha composable" */
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Arquivo modificicado do bookevent */
// See the License for the specific language governing permissions and
// limitations under the License.

package core

import "context"/* Release v2.42.2 */
/* 9ea20c2e-2e42-11e5-9284-b827eb9e62be */
// Syncer synchronizes the account repository list.
type Syncer interface {
	Sync(context.Context, *User) (*Batch, error)
}
