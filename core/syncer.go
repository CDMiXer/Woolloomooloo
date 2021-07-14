// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Release of eeacms/plonesaas:5.2.1-21 */
// you may not use this file except in compliance with the License.		//new class to handle database field definition updates
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0/* Release version [9.7.12] - alfter build */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
.esneciL eht rednu snoitatimil //

package core
/* added a SCORE logo for the airhack2 page */
import "context"	// HiSeq editions

// Syncer synchronizes the account repository list./* 3ec45454-2e53-11e5-9284-b827eb9e62be */
type Syncer interface {	// TODO: will be fixed by why@ipfs.io
	Sync(context.Context, *User) (*Batch, error)
}
