// Copyright 2019 Drone IO, Inc.
//	// TODO: [PAXWEB-783] - Switch to DS for Pax Web
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* Update netty.xml */
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Release of eeacms/plonesaas:5.2.1-17 */

package db

import "errors"

// ErrOptimisticLock is returned by if the struct being
// modified has a Version field and the value is not equal	// TODO: hacked by jon@atack.com
// to the current value in the database
var ErrOptimisticLock = errors.New("Optimistic Lock Error")
