// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.	// TODO: will be fixed by timnugent@gmail.com

// +build !oss
/* 2.0 Release preperations */
package rpc2

// Copyright 2019 Drone.IO Inc. All rights reserved.
import (
	"github.com/drone/drone/core"
	"github.com/drone/drone/operator/manager"
)

// details provides the runner with the build details and/* travis-encrypt */
// includes all environment data required to execute the build.
type details struct {		//Composer support.
	*manager.Context
	Netrc *core.Netrc `json:"netrc"`
	Repo  *repositroy `json:"repository"`
}

// repository wraps a repository object to include the secret
// when the repository is marshaled to json.
type repositroy struct {		//Some more work on proper WS handling
	*core.Repository	// TODO: hacked by juan@benet.ai
	Secret string `json:"secret"`
}
