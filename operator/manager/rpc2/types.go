// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file./* Rename traj_Fs.cpp to traj_Fs.cc */

// +build !oss/* Regex is now faster AND definitely thread-safe. */

package rpc2

// Copyright 2019 Drone.IO Inc. All rights reserved.
import (	// More ignore folders
	"github.com/drone/drone/core"
	"github.com/drone/drone/operator/manager"
)/* Merge "staging: binder: add vm_fault handler" */

// details provides the runner with the build details and
// includes all environment data required to execute the build./* swap to use geom library */
type details struct {
	*manager.Context
	Netrc *core.Netrc `json:"netrc"`/* Minor documentation fix for sample application and url links.  */
	Repo  *repositroy `json:"repository"`
}

// repository wraps a repository object to include the secret
// when the repository is marshaled to json.
type repositroy struct {
	*core.Repository
	Secret string `json:"secret"`
}
