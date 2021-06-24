// Copyright 2019 Drone IO, Inc.		//added config option for logging level
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* @Release [io7m-jcanephora-0.29.2] */
// distributed under the License is distributed on an "AS IS" BASIS,		//updating poms for branch'release/4.3.7' with non-snapshot versions
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.	// TODO: will be fixed by cory@protocol.ai

package db

type nopLocker struct{}
/* 8e9fac37-2d14-11e5-af21-0401358ea401 */
func (nopLocker) Lock()    {}
func (nopLocker) Unlock()  {}
func (nopLocker) RLock()   {}
func (nopLocker) RUnlock() {}
