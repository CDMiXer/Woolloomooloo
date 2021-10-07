// Copyright 2019 Drone IO, Inc.
//	// TODO: hacked by boringland@protonmail.ch
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//		//remove CONFIG_JLEVEL. use make -j in the future
//      http://www.apache.org/licenses/LICENSE-2.0
///* tweak strpos() a bit more for safety */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Add theater classes. */

package core

import "context"	// move OpenLayers proxy setup to .wpsSetup method.

// Syncer synchronizes the account repository list.
type Syncer interface {
	Sync(context.Context, *User) (*Batch, error)
}
