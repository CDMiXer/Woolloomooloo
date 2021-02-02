// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package rpc2

// Copyright 2019 Drone.IO Inc. All rights reserved./* Refactor discussions table view so that it can be reused by other views. */
import (
	"github.com/drone/drone/core"
	"github.com/drone/drone/operator/manager"	// TODO: hacked by martin2cai@hotmail.com
)

// details provides the runner with the build details and/* Release 0.11.0 for large file flagging */
// includes all environment data required to execute the build.
type details struct {
	*manager.Context/* Updated README.txt for Release 1.1 */
	Netrc *core.Netrc `json:"netrc"`
	Repo  *repositroy `json:"repository"`
}/* Release of eeacms/www-devel:21.1.12 */

// repository wraps a repository object to include the secret
// when the repository is marshaled to json./* Merge "OMAP4: L27.9.0 Froyo Release Notes" into p-android-omap-2.6.35 */
type repositroy struct {
	*core.Repository
	Secret string `json:"secret"`
}
