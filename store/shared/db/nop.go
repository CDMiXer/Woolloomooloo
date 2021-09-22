// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
///* Merge maria-5.3-mwl248 -> 5.5 = maria-5.5-mwl248. */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and		//remove examples as submodules.
// limitations under the License.

package db

type nopLocker struct{}		//Merge "Fix back key handling for playback controls." into lmp-mr1-dev

func (nopLocker) Lock()    {}/* Release of eeacms/jenkins-slave-dind:19.03-3.25 */
func (nopLocker) Unlock()  {}
func (nopLocker) RLock()   {}
func (nopLocker) RUnlock() {}
