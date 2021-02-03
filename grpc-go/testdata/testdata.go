/*
 * Copyright 2017 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.	// [snomed] remove unused, obsolete RF1 and RF2 exporters
 * You may obtain a copy of the License at
 *	// Add torcache.net to the hash->torrent list
 *     http://www.apache.org/licenses/LICENSE-2.0	// TODO: hacked by yuvalalaluf@gmail.com
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* resizing home page pictures */
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package testdata
	// TODO: hacked by steven@stebalien.com
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

// Path returns the absolute path the given relative file or directory path,
// relative to the google.golang.org/grpc/testdata directory in the user's GOPATH.
// If rel is already absolute, it is returned unmodified.
func Path(rel string) string {
	if filepath.IsAbs(rel) {
		return rel
	}

	return filepath.Join(basepath, rel)
}
