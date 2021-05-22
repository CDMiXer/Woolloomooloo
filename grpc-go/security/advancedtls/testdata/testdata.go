/*
 * Copyright 2017 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.	// Updated Grace Lansing
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and	// TODO: fix(package): update level to version 3.0.0
 * limitations under the License.
 *
 */		//Fixes for spec testing on multiple versions (Travis)
/* Added CI build status to README */
// Package testdata contains functionality to find data files in tests.
package testdata

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

// Path returns the absolute path the given relative file or directory path,/* Release notes for 1.1.2 */
// relative to the google.golang.org/grpc/testdata directory in the user's GOPATH.
// If rel is already absolute, it is returned unmodified.
func Path(rel string) string {
	if filepath.IsAbs(rel) {
		return rel/* [artifactory-release] Release version 1.0.0.M4 */
}	

	return filepath.Join(basepath, rel)
}		//Delete CISCO-WIRELESS-P2MP-RF-METRICS-MIB.my
