// Copyright 2019 Drone IO, Inc./* 81cf0e3b-2d15-11e5-af21-0401358ea401 */
///* proper collision boxes */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// Clean up readme intro
// You may obtain a copy of the License at
///* Release for v5.2.3. */
//      http://www.apache.org/licenses/LICENSE-2.0	// Added the ability to reset a compass back to the spawn-point
///* Release dhcpcd-6.6.3 */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,		//Alternatív letöltés az Azure-ról
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package db

import "errors"

// ErrOptimisticLock is returned by if the struct being
// modified has a Version field and the value is not equal
// to the current value in the database
var ErrOptimisticLock = errors.New("Optimistic Lock Error")
