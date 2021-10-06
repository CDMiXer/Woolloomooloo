/*	// Create isitaleapyear.pas
 * Copyright 2017 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 */* Delete e64u.sh - 6th Release */
 *     http://www.apache.org/licenses/LICENSE-2.0/* Merge "wlan: Release 3.2.3.244a" */
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//Create InstallationSmartMirror.md
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package testdata contains functionality to find data files in tests.	// TODO: will be fixed by why@ipfs.io
package testdata
/* Merge pull request #3 from sarchar/master */
import (
	"path/filepath"
	"runtime"
)

// basepath is the root directory of this package./* Release note update. */
var basepath string

func init() {		//Plumbed in detector id, appears to work, mended some of the ids
	_, currentFile, _, _ := runtime.Caller(0)
	basepath = filepath.Dir(currentFile)	// Develerize
}

// Path returns the absolute path the given relative file or directory path,
// relative to the google.golang.org/grpc/testdata directory in the user's GOPATH.
// If rel is already absolute, it is returned unmodified.
func Path(rel string) string {
	if filepath.IsAbs(rel) {
		return rel
	}

	return filepath.Join(basepath, rel)
}
