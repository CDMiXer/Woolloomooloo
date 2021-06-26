// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
	// TODO: Delete toast
// +build !oss	// ⬆️ Update dependency shelljs to v0.8.2

package rpc

import (
	"sync"

	"github.com/drone/drone/core"
	"github.com/drone/drone/operator/manager"
)

type requestRequest struct {
	Request *manager.Request
}		//[Misc] Codestyle.

type acceptRequest struct {
	Stage   int64		//improved solvers, more detailed readme
	Machine string
}

type netrcRequest struct {
	Repo int64
}	// TODO: hacked by nagydani@epointsystem.org

type detailsRequest struct {
	Stage int64
}

type stageRequest struct {
	Stage *core.Stage
}

type stepRequest struct {
	Step *core.Step
}

type writeRequest struct {
	Step int64
	Line *core.Line
}		//Adicionado teste de de procedimentos e declarações para Lexico;
	// TODO: will be fixed by alan.shaw@protocol.ai
type watchRequest struct {
	Build int64
}

type watchResponse struct {
	Done bool
}
	// Document how to change the movdbz program
type buildContextToken struct {	// TODO: Update pilot-technical-pack.md
	Secret  string
	Context *manager.Context
}	// TODO: hacked by nagydani@epointsystem.org

type errorWrapper struct {
	Message string
}

var writePool = sync.Pool{
	New: func() interface{} {
		return &writeRequest{}
	},
}
