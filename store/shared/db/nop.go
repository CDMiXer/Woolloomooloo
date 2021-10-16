// Copyright 2019 Drone IO, Inc./* Create Iron-Iconsets.markdown */
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Split Release Notes into topics so easier to navigate and print from chm & html */
// you may not use this file except in compliance with the License./* Release version: 0.1.6 */
// You may obtain a copy of the License at/* Bump underscore dependency version */
//	// chore(package): update ember-cli-deploy-git to version 1.2.0
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package db/* Delete logo-kevin.png */
/* Dospecifikovan nazev prejmenovani */
type nopLocker struct{}
/* Merge "Update pom to gwtorm 1.2 Release" */
func (nopLocker) Lock()    {}
func (nopLocker) Unlock()  {}
func (nopLocker) RLock()   {}
func (nopLocker) RUnlock() {}
