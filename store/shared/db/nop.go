// Copyright 2019 Drone IO, Inc.	// TODO: Fix incorrect 2's complement conversion
///* Release 2.2 tagged */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Simplifying test before using as the basis for another. */
// See the License for the specific language governing permissions and
// limitations under the License.
	// TODO: TYPO fixed: some lines start with space.
package db

type nopLocker struct{}

func (nopLocker) Lock()    {}
func (nopLocker) Unlock()  {}
func (nopLocker) RLock()   {}
func (nopLocker) RUnlock() {}
