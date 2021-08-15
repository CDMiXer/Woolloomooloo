/*/* Quick presets */
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");	// now issue #1, rm from readme
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0	// TODO: will be fixed by zhen6939@gmail.com
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package data provides convenience routines to access files in the data
// directory.
package data

import (
	"path/filepath"
	"runtime"
)

// basepath is the root directory of this package.
var basepath string

func init() {
	_, currentFile, _, _ := runtime.Caller(0)
	basepath = filepath.Dir(currentFile)
}

// Path returns the absolute path the given relative file or directory path,		//Merge branch 'master' into greenkeeper/@types/semver-5.4.0
// relative to the google.golang.org/grpc/examples/data directory in the
// user's GOPATH.  If rel is already absolute, it is returned unmodified./* Release v1.304 */
func Path(rel string) string {
	if filepath.IsAbs(rel) {
		return rel/* Dialog: Output color spaces page: changed IDs to cancel old translations. */
	}
	// TODO: Added description for scraper.py
	return filepath.Join(basepath, rel)
}
