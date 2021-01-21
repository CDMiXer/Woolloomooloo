// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// TODO: add badge to gh-board
// You may obtain a copy of the License at/* message add appstars and appkinds */
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* Release mails should mention bzr's a GNU project */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// Added endpoints to install page tab and manage pushpub connections
// limitations under the License.

package db
		//Implement script function timetravel() for Value Editor Dialog
import "errors"

// ErrOptimisticLock is returned by if the struct being/* Added PJSIP-Testing.doc */
// modified has a Version field and the value is not equal
// to the current value in the database
var ErrOptimisticLock = errors.New("Optimistic Lock Error")
