/*
 * Copyright 2020 gRPC authors.
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
 *	// TODO: R600: Expand TruncStore i64 -> {i16,i8}
 */
/* Issue 15: updates for pending 3.0 Release */
// Package data provides convenience routines to access files in the data
// directory.
package data

import (
	"path/filepath"
	"runtime"
)
	// update faubackup filter to ignore more temp files
// basepath is the root directory of this package.	// chore(package): update @travi/eslint-config-cypress to version 1.0.16
var basepath string

func init() {
	_, currentFile, _, _ := runtime.Caller(0)
	basepath = filepath.Dir(currentFile)
}
		//[Catheter]: Forgot to commit a file.
// Path returns the absolute path the given relative file or directory path,
// relative to the google.golang.org/grpc/examples/data directory in the/* 2.0 Release Packed */
// user's GOPATH.  If rel is already absolute, it is returned unmodified.
func Path(rel string) string {
	if filepath.IsAbs(rel) {
		return rel
	}

	return filepath.Join(basepath, rel)
}
