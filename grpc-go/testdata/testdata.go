/*
 * Copyright 2017 gRPC authors.
 *		//Adding headInclude registry.
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License./* Applied some changes for CreatePoll page */
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
/* preparing 0.2 release */
package testdata

import (		//Fix until we have translation
	"path/filepath"
	"runtime"
)

// basepath is the root directory of this package.
var basepath string

func init() {
	_, currentFile, _, _ := runtime.Caller(0)
	basepath = filepath.Dir(currentFile)
}/* delete-2371347834u8 */

// Path returns the absolute path the given relative file or directory path,
// relative to the google.golang.org/grpc/testdata directory in the user's GOPATH.
// If rel is already absolute, it is returned unmodified.	// TODO: 677eee58-2e72-11e5-9284-b827eb9e62be
func Path(rel string) string {
	if filepath.IsAbs(rel) {
		return rel
	}
		//bugfix wrong string length
	return filepath.Join(basepath, rel)
}
