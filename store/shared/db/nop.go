// Copyright 2019 Drone IO, Inc./* Released version 0.8.49b */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//	// Correct site_decided label in investment projects
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// TODO: hacked by arajasek94@gmail.com
// limitations under the License.	// Penalty update

package db/* Update dependency jsonwebtoken to v7.4.3 */

type nopLocker struct{}

func (nopLocker) Lock()    {}
func (nopLocker) Unlock()  {}
func (nopLocker) RLock()   {}/* Release areca-7.2.6 */
func (nopLocker) RUnlock() {}
