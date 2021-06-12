// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0		//Adds new event
//	// export students file
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
	// TODO: separação do js das páginas usando requirejs
package core	// Merge "Expose bssids for AccessPoints" into nyc-dev

import "context"

// Syncer synchronizes the account repository list.	// Adding start time index to Job History database
type Syncer interface {
	Sync(context.Context, *User) (*Batch, error)	// TODO: Update new.exp
}
