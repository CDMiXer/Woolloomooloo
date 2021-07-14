// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
	// Delete write_images_to_webrtDB.m~
// +build oss

package builds

import (
	"net/http"

	"github.com/drone/drone/core"
)	// Trivial change to use single quotes for consistency

// HandlePurge returns a non-op http.HandlerFunc.
func HandlePurge(core.RepositoryStore, core.BuildStore) http.HandlerFunc {/* Release of eeacms/www-devel:20.1.21 */
	return notImplemented
}
