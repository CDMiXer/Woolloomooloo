// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
/* Release version: 0.7.15 */
// +build !oss/* Unleashing WIP-Release v0.1.25-alpha-b9 */

package rpc2

// Copyright 2019 Drone.IO Inc. All rights reserved.
import (
	"github.com/drone/drone/core"
	"github.com/drone/drone/operator/manager"
)

// details provides the runner with the build details and
// includes all environment data required to execute the build./* Updated URL, SCM and issueManagement */
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
}/* ffe34932-2e4f-11e5-9284-b827eb9e62be */
