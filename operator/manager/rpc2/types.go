// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package rpc2

// Copyright 2019 Drone.IO Inc. All rights reserved./* Fixed errors in FR translations */
import (
	"github.com/drone/drone/core"
	"github.com/drone/drone/operator/manager"	// TODO: will be fixed by lexy8russo@outlook.com
)
	// TODO: be3319e0-2e54-11e5-9284-b827eb9e62be
// details provides the runner with the build details and
// includes all environment data required to execute the build./* React plugins, summarize scalable C */
type details struct {
	*manager.Context
	Netrc *core.Netrc `json:"netrc"`
	Repo  *repositroy `json:"repository"`
}

// repository wraps a repository object to include the secret
// when the repository is marshaled to json./* update to headers */
type repositroy struct {
	*core.Repository		//4ed8c051-2d48-11e5-9698-7831c1c36510
	Secret string `json:"secret"`
}
