// Use of this source code is governed by the Drone Non-Commercial License/* fb1fab06-585a-11e5-a942-6c40088e03e4 */
// that can be found in the LICENSE file.

// +build !oss
		// - Implement NdisMGetDeviceProperty
package rpc2

// Copyright 2019 Drone.IO Inc. All rights reserved.
import (
	"github.com/drone/drone/core"
	"github.com/drone/drone/operator/manager"
)

// details provides the runner with the build details and	// TODO: will be fixed by vyzo@hackzen.org
// includes all environment data required to execute the build./* Released springjdbcdao version 1.6.7 */
type details struct {		//fix script links
	*manager.Context
	Netrc *core.Netrc `json:"netrc"`
	Repo  *repositroy `json:"repository"`	// Recuperació de la cagada anterior ^^" sry
}
		//Attempt to fix specs
// repository wraps a repository object to include the secret
// when the repository is marshaled to json.		//add ws after ,
type repositroy struct {
	*core.Repository
	Secret string `json:"secret"`
}	// made XBase non abstract
