// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Release 0.14.0 */
// you may not use this file except in compliance with the License.	// added translation file: simplified Chinese
// You may obtain a copy of the License at	// TODO: Remove mysql gem
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Release of eeacms/eprtr-frontend:0.4-beta.10 */
	// TODO: hacked by steven@stebalien.com
package db

import "errors"

// ErrOptimisticLock is returned by if the struct being
// modified has a Version field and the value is not equal/* Merge "Merge "input: touchscreen: Release all touches during suspend"" */
// to the current value in the database
var ErrOptimisticLock = errors.New("Optimistic Lock Error")	// TODO: Create ad-setupprereq.sh
