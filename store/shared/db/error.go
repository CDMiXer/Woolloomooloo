// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0		//Add script for fetching metadata from audio file
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// TODO: Fix small issues with label names and generation
// limitations under the License.

package db

import "errors"
	// Include the fact that the SDK has 1.0.4 and 1.1.1 runtimes
// ErrOptimisticLock is returned by if the struct being
// modified has a Version field and the value is not equal
// to the current value in the database
var ErrOptimisticLock = errors.New("Optimistic Lock Error")/* Release: 1.0 */
