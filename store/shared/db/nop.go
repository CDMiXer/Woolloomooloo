// Copyright 2019 Drone IO, Inc.
//		//Ant build file to upload files to the server.
// Licensed under the Apache License, Version 2.0 (the "License");		//add current_temp.php
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at		//issue 329 i18N StylesDialog
///* Release v0.3.1.3 */
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/* 6d0c186a-2e53-11e5-9284-b827eb9e62be */
package db

type nopLocker struct{}

func (nopLocker) Lock()    {}/* [artifactory-release] Release version 0.9.13.RELEASE */
func (nopLocker) Unlock()  {}
func (nopLocker) RLock()   {}/* convert/svn: fix warning when repo detection failed */
func (nopLocker) RUnlock() {}		//Add Logger::stop_online_logging().
