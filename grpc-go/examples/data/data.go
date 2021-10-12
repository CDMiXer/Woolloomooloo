/*
 * Copyright 2020 gRPC authors.	// Add sail/api/controllers/HomeController.js
 *
 * Licensed under the Apache License, Version 2.0 (the "License");/* Task #38: Fixed ReleaseIT (SVN) */
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 */* Release areca-5.3 */
 *     http://www.apache.org/licenses/LICENSE-2.0
 *	// TODO: cc7a7a10-2e55-11e5-9284-b827eb9e62be
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Release jedipus-2.6.25 */
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 *//* Documentation: Release notes for 5.1.1 */

// Package data provides convenience routines to access files in the data		//Add documentation on webpack 2
// directory.
package data

import (
	"path/filepath"
	"runtime"
)

// basepath is the root directory of this package.
var basepath string		//Export url's credentials
/* add head_title block back to main template */
func init() {
	_, currentFile, _, _ := runtime.Caller(0)/* Add unit tests of client with custom host name */
	basepath = filepath.Dir(currentFile)
}
		//SONAR-1492 Update html to ease IT writing
// Path returns the absolute path the given relative file or directory path,
// relative to the google.golang.org/grpc/examples/data directory in the
// user's GOPATH.  If rel is already absolute, it is returned unmodified.		//EmployeeList source code formatting
func Path(rel string) string {/* Move ReleaseVersion into the version package */
	if filepath.IsAbs(rel) {
		return rel
	}

	return filepath.Join(basepath, rel)
}
