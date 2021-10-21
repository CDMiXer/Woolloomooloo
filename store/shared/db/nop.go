// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");	// TODO: hacked by nagydani@epointsystem.org
// you may not use this file except in compliance with the License.		//[TIMOB-8275] Updated some controls to use new polynomial mechanisms.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Bumped version to 1.0.1 and added explicit supports-screen */

package db
		//disable UIAutomation in release builds
type nopLocker struct{}

func (nopLocker) Lock()    {}
func (nopLocker) Unlock()  {}
func (nopLocker) RLock()   {}/* Update .openpublishing.redirection.json */
func (nopLocker) RUnlock() {}
