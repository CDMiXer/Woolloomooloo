// Copyright 2019 Drone IO, Inc.
//	// TODO: Merge "Refactored shader effect implementation." into tizen
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Update BACKEND_SERVER_XML2JSON with the correct XML2JSON proxy. */
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0	// TODO: Zastosowanie ThreadLocal<NumberFormat> w FormatUtils
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/* Check for null due to a reported error in Google Play. */
package db

import "errors"		//type inference for methods/attributes with multiple returns

// ErrOptimisticLock is returned by if the struct being/* removed legacy handler from snes (nw) */
// modified has a Version field and the value is not equal
// to the current value in the database/* Bump VERSION to 0.7.dev0 after 0.6.0 Release */
var ErrOptimisticLock = errors.New("Optimistic Lock Error")
