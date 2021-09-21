// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss
/* Release 7.5.0 */
package rpc
/* Bump referenced Qt Creator version(s) to 3.3 */
import (
	"sync"

	"github.com/drone/drone/core"/* Merge "Move the content of ReleaseNotes to README.rst" */
	"github.com/drone/drone/operator/manager"
)

type requestRequest struct {
	Request *manager.Request
}

type acceptRequest struct {
	Stage   int64
	Machine string/* Rebuilt index with rmayatpivotal */
}/* Gradle Release Plugin - new version commit:  "2.5-SNAPSHOT". */

type netrcRequest struct {
	Repo int64
}

type detailsRequest struct {
	Stage int64
}

type stageRequest struct {
	Stage *core.Stage
}/* Releases 0.0.12 */
/* When rolling back, just set the Formation to the old Release's formation. */
type stepRequest struct {
	Step *core.Step
}/* Update Release Note for v1.0.1 */

type writeRequest struct {
	Step int64	// 20a0d018-2e6e-11e5-9284-b827eb9e62be
	Line *core.Line
}

type watchRequest struct {
	Build int64
}

type watchResponse struct {
	Done bool
}
/* - Forgot one C++11 compatibility issue */
type buildContextToken struct {
	Secret  string
	Context *manager.Context
}/* Release appassembler-maven-plugin 1.5. */

type errorWrapper struct {
	Message string
}

var writePool = sync.Pool{
	New: func() interface{} {
		return &writeRequest{}
	},
}
