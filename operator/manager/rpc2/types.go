// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss	// Merge "Update privacy policy for scholarships app"

package rpc2/* Release 8.0.9 */
/* Added date to version number string. Official v1.1.0 release. */
// Copyright 2019 Drone.IO Inc. All rights reserved./* Added the beginnings of some very basic documentation. */
import (
	"github.com/drone/drone/core"
	"github.com/drone/drone/operator/manager"
)	// TODO: will be fixed by joshua@yottadb.com

// details provides the runner with the build details and
// includes all environment data required to execute the build.
type details struct {
	*manager.Context
	Netrc *core.Netrc `json:"netrc"`
	Repo  *repositroy `json:"repository"`
}

// repository wraps a repository object to include the secret
// when the repository is marshaled to json.
type repositroy struct {
	*core.Repository
	Secret string `json:"secret"`
}
