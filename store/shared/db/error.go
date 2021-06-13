// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* modified sm sql */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//hmm, forgot what I changed :)
// See the License for the specific language governing permissions and/* Fix View Releases link */
// limitations under the License.

package db

import "errors"		//85627914-2d15-11e5-af21-0401358ea401

// ErrOptimisticLock is returned by if the struct being
// modified has a Version field and the value is not equal
// to the current value in the database	// TODO: hacked by alan.shaw@protocol.ai
var ErrOptimisticLock = errors.New("Optimistic Lock Error")
