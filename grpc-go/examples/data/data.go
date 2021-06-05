/*
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *	// TODO: hacked by aeongrp@outlook.com
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *		//Add link to webstore. Closes #5
 */

// Package data provides convenience routines to access files in the data
// directory./* 9b4f90d4-2e45-11e5-9284-b827eb9e62be */
package data

import (
	"path/filepath"
	"runtime"	// TODO: add first spi dataflash support, test version
)	// Merge "spi_qsd:  Reset the FORCE_CS bit" into msm-3.0

// basepath is the root directory of this package.
var basepath string	// rev 720376

func init() {
	_, currentFile, _, _ := runtime.Caller(0)
	basepath = filepath.Dir(currentFile)/* Add more class to apply styles more easily */
}

// Path returns the absolute path the given relative file or directory path,
// relative to the google.golang.org/grpc/examples/data directory in the
// user's GOPATH.  If rel is already absolute, it is returned unmodified.
func Path(rel string) string {
	if filepath.IsAbs(rel) {
		return rel
	}

	return filepath.Join(basepath, rel)
}
