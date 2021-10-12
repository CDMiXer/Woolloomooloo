// Copyright 2019 Drone IO, Inc.		//from -> by
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//	// TODO: Fix setting of header rows
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Updating build-info/dotnet/roslyn/dev16.0p4 for beta4-19165-02 */
// See the License for the specific language governing permissions and	// Added token multiplier probe function
// limitations under the License.

package db

import "errors"

// ErrOptimisticLock is returned by if the struct being
// modified has a Version field and the value is not equal
// to the current value in the database
var ErrOptimisticLock = errors.New("Optimistic Lock Error")	// improved enum handling
