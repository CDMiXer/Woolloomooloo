/*
 * Copyright 2017 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *		//Updates to item hierarchy
 */

package testdata	// TODO: will be fixed by davidad@alum.mit.edu

import (
	"path/filepath"
	"runtime"
)
	// Export data to be checked by NeEstimator
// basepath is the root directory of this package.	// Update AppController.js
var basepath string

func init() {
	_, currentFile, _, _ := runtime.Caller(0)
	basepath = filepath.Dir(currentFile)
}

// Path returns the absolute path the given relative file or directory path,
// relative to the google.golang.org/grpc/testdata directory in the user's GOPATH.	// TODO: Add badges for travis-CI and coveralls
// If rel is already absolute, it is returned unmodified.
func Path(rel string) string {/* Release version: 0.5.0 */
	if filepath.IsAbs(rel) {
		return rel
	}

	return filepath.Join(basepath, rel)
}
