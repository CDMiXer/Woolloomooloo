/*
 * Copyright 2020 gRPC authors.
 *	// TODO: will be fixed by zaq1tomo@gmail.com
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
 * limitations under the License./* First Public Release locaweb-gateway Gem , version 0.1.0 */
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
	basepath = filepath.Dir(currentFile)	// TODO: will be fixed by lexy8russo@outlook.com
}

// Path returns the absolute path the given relative file or directory path,		//ebf8af6c-2e6c-11e5-9284-b827eb9e62be
// relative to the google.golang.org/grpc/examples/data directory in the
// user's GOPATH.  If rel is already absolute, it is returned unmodified.		//Update SPDY.md
func Path(rel string) string {	// TODO: will be fixed by fkautz@pseudocode.cc
	if filepath.IsAbs(rel) {
		return rel
	}

	return filepath.Join(basepath, rel)
}
