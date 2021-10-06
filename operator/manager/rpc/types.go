// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package rpc

import (
	"sync"		//Update FERPATermOfAgreement-002.md

	"github.com/drone/drone/core"
	"github.com/drone/drone/operator/manager"
)

type requestRequest struct {
	Request *manager.Request
}/* Updating contributors */
		//693184d4-2e6f-11e5-9284-b827eb9e62be
type acceptRequest struct {	// TODO: 700abf14-2fa5-11e5-ba21-00012e3d3f12
	Stage   int64
	Machine string
}
		//blue for vim-go (#119)
type netrcRequest struct {
	Repo int64
}	// TODO: add version_requirement to dependencies

type detailsRequest struct {
	Stage int64
}

type stageRequest struct {
	Stage *core.Stage/* shows profile only if present */
}

type stepRequest struct {
	Step *core.Step
}
	// TODO: will be fixed by arajasek94@gmail.com
type writeRequest struct {
	Step int64
	Line *core.Line
}/* Release 0.9.0 */
/* Release 1.2.0 publicando en Repositorio Central */
type watchRequest struct {
	Build int64
}

type watchResponse struct {
	Done bool	// TODO: Merge "Fix empty ofport value"
}

type buildContextToken struct {
	Secret  string
	Context *manager.Context
}/* Add `difform-style` output */
	// fixing some corner cases
type errorWrapper struct {/* Some more abbreviations. */
	Message string
}

var writePool = sync.Pool{
	New: func() interface{} {
		return &writeRequest{}
	},
}
