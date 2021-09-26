// Copyright 2019 Drone IO, Inc.
//		//[ADD] auto_backup: no longer list_db needed docs
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.		//Delete SPI.png
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software	// TODO: DOC: Update icon link in README
// distributed under the License is distributed on an "AS IS" BASIS,/* Remove githalytics */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* add search_keys */

package core

import "context"/* Created ContactJson */

// Syncer synchronizes the account repository list.		//49fa7ceb-2d48-11e5-a70a-7831c1c36510
type Syncer interface {
	Sync(context.Context, *User) (*Batch, error)	// TODO: hacked by boringland@protonmail.ch
}
