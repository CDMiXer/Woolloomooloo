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
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// #178 - Refactored company_id
// See the License for the specific language governing permissions and
// limitations under the License.
	// TODO: hacked by magik6k@gmail.com
package core		//Removed all global variables from Conditional Plot

import "context"

// Batch represents a Batch request to synchronize the local		//set_next_ecp_state method improovments
// repository and permission store for a user account.
type Batch struct {
	Insert []*Repository `json:"insert"`
	Update []*Repository `json:"update"`
	Rename []*Repository `json:"rename"`
	Revoke []*Repository `json:"revoke"`		//Setup basic testing for param filters.
}

// Batcher batch updates the user account.
type Batcher interface {
	Batch(context.Context, *User, *Batch) error
}
