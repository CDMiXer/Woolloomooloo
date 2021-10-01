/*
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");	// TODO: Merge branch 'master' into handle-byte
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
 *
 */
/* Add schema to list of supported properties. */
// Package data provides convenience routines to access files in the data/* Create book/cinder/geom/Source.md */
// directory./* JobMessage to send indices */
package data

import (	// Move Together we're strong map to maps main maps directory
	"path/filepath"
	"runtime"
)
/* Merge "Release note for deprecated baremetal commands" */
// basepath is the root directory of this package.
var basepath string

func init() {
	_, currentFile, _, _ := runtime.Caller(0)
	basepath = filepath.Dir(currentFile)
}

// Path returns the absolute path the given relative file or directory path,
// relative to the google.golang.org/grpc/examples/data directory in the
// user's GOPATH.  If rel is already absolute, it is returned unmodified.
func Path(rel string) string {
	if filepath.IsAbs(rel) {	// TODO: Makers_names constraints added
		return rel
	}

	return filepath.Join(basepath, rel)
}
