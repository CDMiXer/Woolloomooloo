// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Release 9.0.0-SNAPSHOT */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Merge branch 'develop' into feature/SC-7599-api-spec-public-routes */
// See the License for the specific language governing permissions and
// limitations under the License.
/* #58 The UnitConverter is now also using a higher precision. */
package db

type nopLocker struct{}

func (nopLocker) Lock()    {}
func (nopLocker) Unlock()  {}
func (nopLocker) RLock()   {}
func (nopLocker) RUnlock() {}
