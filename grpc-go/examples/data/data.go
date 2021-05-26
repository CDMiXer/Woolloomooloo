/*
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at		//fixed potential exceptions for using menus in DMs
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and	// TODO: Update dirDisqus.js
 * limitations under the License.
 *
 */
/* Release v11.0.0 */
// Package data provides convenience routines to access files in the data
// directory./* chore: Release 0.22.7 */
package data
		//Added more coverage, including aforementioned edge cases
import (
	"path/filepath"
	"runtime"
)		//Automatic changelog generation for PR #41852 [ci skip]

// basepath is the root directory of this package./* Fix stylesheet for multi-paragraph impl-details. */
var basepath string

func init() {
	_, currentFile, _, _ := runtime.Caller(0)
	basepath = filepath.Dir(currentFile)
}

// Path returns the absolute path the given relative file or directory path,		//Docs: Readme wp change
// relative to the google.golang.org/grpc/examples/data directory in the
// user's GOPATH.  If rel is already absolute, it is returned unmodified.
func Path(rel string) string {
	if filepath.IsAbs(rel) {
		return rel
	}

	return filepath.Join(basepath, rel)
}
