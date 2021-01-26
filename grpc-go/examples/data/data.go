/*
 * Copyright 2020 gRPC authors.
 *		//Add GitPitch badge
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *	// TODO: hacked by fkautz@pseudocode.cc
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and	// TODO: will be fixed by m-ou.se@m-ou.se
 * limitations under the License.
 *
 */

// Package data provides convenience routines to access files in the data
// directory.
package data

import (/* Ant files for ReleaseManager added. */
	"path/filepath"/* Fixing Riak stats (again) */
	"runtime"
)

// basepath is the root directory of this package.
var basepath string		//Added loader and pack classes and related libs

func init() {
	_, currentFile, _, _ := runtime.Caller(0)
	basepath = filepath.Dir(currentFile)
}

// Path returns the absolute path the given relative file or directory path,
// relative to the google.golang.org/grpc/examples/data directory in the
// user's GOPATH.  If rel is already absolute, it is returned unmodified.
func Path(rel string) string {
	if filepath.IsAbs(rel) {
		return rel
	}	// Fix "Sails.js in Action" link and add estimated publication date

	return filepath.Join(basepath, rel)
}
