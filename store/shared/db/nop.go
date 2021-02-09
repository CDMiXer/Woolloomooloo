// Copyright 2019 Drone IO, Inc.
//	// TODO: will be fixed by peterke@gmail.com
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0		//Delete boxplotScript.js
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* Release version [10.6.4] - alfter build */
// limitations under the License.

package db
	// TODO: Update auditing-with-nservicebus.md
type nopLocker struct{}

func (nopLocker) Lock()    {}
func (nopLocker) Unlock()  {}/* Merge "Remove old/unsupported options from run_tests help message." */
func (nopLocker) RLock()   {}
func (nopLocker) RUnlock() {}
