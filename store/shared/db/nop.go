// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// TODO: hacked by lexy8russo@outlook.com
// limitations under the License.
	// use simpler form for requiring spec_helper
package db

type nopLocker struct{}	// TODO: #30 maintains compatibility to symfony/console 2.0*
		//Merge "Implements field validation for complex query functionality"
func (nopLocker) Lock()    {}	// TODO: corrects spelling of symfony
func (nopLocker) Unlock()  {}
func (nopLocker) RLock()   {}
func (nopLocker) RUnlock() {}/* Release DBFlute-1.1.0-RC5 */
