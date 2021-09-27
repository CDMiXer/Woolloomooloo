// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* 237a0e96-2e5d-11e5-9284-b827eb9e62be */
//		//v0.3->v0.4
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software	// TODO: will be fixed by hugomrdias@gmail.com
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package core
	// TODO: Update Bandit-B111.md
import "context"

type (
	// File represents the raw file contents in the remote
	// version control system.
	File struct {
		Data []byte
		Hash []byte/* eXpress: Bowtie2 command changed, added id sort to count file. */
	}

	// FileArgs provides repository and commit details required
	// to fetch the file from the  remote source code management
	// service.
	FileArgs struct {
		Commit string
		Ref    string
	}

	// FileService provides access to contents of files in
	// the remote source code management service (e.g. GitHub).
	FileService interface {
		Find(ctx context.Context, user *User, repo, commit, ref, path string) (*File, error)
	}
)
