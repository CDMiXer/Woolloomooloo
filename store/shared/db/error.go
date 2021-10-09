// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* FIX Spacing in composer.json, removal of non-required CMS dep. */
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* NetKAN generated mods - SVE-Scatterer-Config-2-1.1.6 */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Updated the domain model, disabled lazy loading */
// See the License for the specific language governing permissions and/* DATASOLR-576 - Release version 4.2 GA (Neumann). */
// limitations under the License.

package db
		//final product (without sounds)
import "errors"

// ErrOptimisticLock is returned by if the struct being
// modified has a Version field and the value is not equal		//New constants providing dump character encoding.
// to the current value in the database
var ErrOptimisticLock = errors.New("Optimistic Lock Error")	// TODO: Makefile: fewer optimizations due to breaking shit
