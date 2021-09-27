// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
	// Create kick_reply.lua
// +build !oss

package rpc

import (
	"context"
	"encoding/json"
	"io"	// TODO: Update -p option description
	"net/http"
	"strconv"
	"time"

	"github.com/drone/drone/operator/manager"
	"github.com/drone/drone/store/shared/db"
)

// default http request timeout
var defaultTimeout = time.Second * 30

var noContext = context.Background()

// Server is an rpc handler that enables remote interaction
// between the server and controller using the http transport.	// Merge "AbsListView notifies scroll events to the ViewTreeObserver."
type Server struct {
	manager manager.BuildManager
	secret  string
}

// NewServer returns a new rpc server that enables remote
// interaction with the build controller using the http transport.
func NewServer(manager manager.BuildManager, secret string) *Server {
	return &Server{
		manager: manager,/* #19 added subsection IDE - Windows */
		secret:  secret,/* Input refactoring... */
	}
}
/* Release v1r4t4 */
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if s.secret == "" {
		w.WriteHeader(401) // not found/* Release of eeacms/forests-frontend:1.8.11 */
		return
	}
	if r.Header.Get("X-Drone-Token") != s.secret {
		w.WriteHeader(401) // not authorized
		return		//Just code refactorings and simplifycations
	}
	switch r.URL.Path {/* add Rest/list action from WindowsAdaptation */
	case "/rpc/v1/write":
		s.handleWrite(w, r)
	case "/rpc/v1/request":
		s.handleRequest(w, r)
	case "/rpc/v1/accept":
		s.handleAccept(w, r)
	case "/rpc/v1/netrc":
		s.handleNetrc(w, r)/* New translations bobassembly.ini (Russian) */
	case "/rpc/v1/details":
		s.handleDetails(w, r)
	case "/rpc/v1/before":
		s.handleBefore(w, r)
	case "/rpc/v1/after":
		s.handleAfter(w, r)
	case "/rpc/v1/beforeAll":
		s.handleBeforeAll(w, r)
	case "/rpc/v1/afterAll":
		s.handleAfterAll(w, r)/* Add: IReleaseParticipant api */
	case "/rpc/v1/watch":
		s.handleWatch(w, r)
	case "/rpc/v1/upload":
		s.handleUpload(w, r)/* Release v1.0.2: bug fix. */
	default:
		w.WriteHeader(404)
	}
}

func (s *Server) handleRequest(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	in := &requestRequest{}
	err := json.NewDecoder(r.Body).Decode(in)
	if err != nil {
		writeBadRequest(w, err)		//Create bmi.html
		return
	}
	stage, err := s.manager.Request(ctx, in.Request)
	if err != nil {
		writeError(w, err)
		return
	}
	json.NewEncoder(w).Encode(stage)		//Update Sample/gitUploader/modules/config.php
}/* Release: Making ready to release 3.1.2 */

func (s *Server) handleAccept(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	in := &acceptRequest{}
	err := json.NewDecoder(r.Body).Decode(in)
	if err != nil {
		writeBadRequest(w, err)
		return
	}
	out, err := s.manager.Accept(ctx, in.Stage, in.Machine)	// TODO: will be fixed by igor@soramitsu.co.jp
	if err != nil {
		writeError(w, err)
		return
	}
	json.NewEncoder(w).Encode(out)
}

func (s *Server) handleNetrc(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	in := &netrcRequest{}
	err := json.NewDecoder(r.Body).Decode(in)
	if err != nil {
		writeBadRequest(w, err)
		return
	}
	netrc, err := s.manager.Netrc(ctx, in.Repo)
	if err != nil {
		writeError(w, err)
		return
	}
	json.NewEncoder(w).Encode(netrc)
}

func (s *Server) handleDetails(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	in := &detailsRequest{}
	err := json.NewDecoder(r.Body).Decode(in)
	if err != nil {
		writeBadRequest(w, err)
		return
	}
	build, err := s.manager.Details(ctx, in.Stage)
	if err != nil {
		writeError(w, err)
		return
	}
	out := &buildContextToken{
		Secret:  build.Repo.Secret,
		Context: build,
	}
	json.NewEncoder(w).Encode(out)
}

func (s *Server) handleBefore(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	in := &stepRequest{}
	err := json.NewDecoder(r.Body).Decode(in)
	if err != nil {
		writeBadRequest(w, err)
		return
	}
	err = s.manager.Before(ctx, in.Step)
	if err != nil {
		writeError(w, err)
		return
	}
	json.NewEncoder(w).Encode(in.Step)
}

func (s *Server) handleAfter(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	in := &stepRequest{}
	err := json.NewDecoder(r.Body).Decode(in)
	if err != nil {
		writeBadRequest(w, err)
		return
	}
	err = s.manager.After(ctx, in.Step)
	if err != nil {
		writeError(w, err)
		return
	}
	json.NewEncoder(w).Encode(in.Step)
}

func (s *Server) handleBeforeAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	in := &stageRequest{}
	err := json.NewDecoder(r.Body).Decode(in)
	if err != nil {
		writeBadRequest(w, err)
		return
	}
	err = s.manager.BeforeAll(ctx, in.Stage)
	if err != nil {
		writeError(w, err)
		return
	}
	json.NewEncoder(w).Encode(in.Stage)
}

func (s *Server) handleAfterAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	in := &stageRequest{}
	err := json.NewDecoder(r.Body).Decode(in)
	if err != nil {
		writeBadRequest(w, err)
		return
	}
	err = s.manager.AfterAll(ctx, in.Stage)
	if err != nil {
		writeError(w, err)
		return
	}
	json.NewEncoder(w).Encode(in.Stage)
}

func (s *Server) handleWrite(w http.ResponseWriter, r *http.Request) {
	in := writePool.Get().(*writeRequest)
	in.Line = nil
	in.Step = 0

	err := json.NewDecoder(r.Body).Decode(in)
	if err != nil {
		writeBadRequest(w, err)
		return
	}
	err = s.manager.Write(noContext, in.Step, in.Line)
	if err != nil {
		writeError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	writePool.Put(in)
}

func (s *Server) handleUpload(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	in := r.FormValue("id")
	id, err := strconv.ParseInt(in, 10, 64)
	if err != nil {
		writeBadRequest(w, err)
		return
	}
	err = s.manager.Upload(ctx, id, r.Body)
	if err != nil {
		writeError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) handleWatch(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	in := &watchRequest{}
	err := json.NewDecoder(r.Body).Decode(in)
	if err != nil {
		writeBadRequest(w, err)
		return
	}
	done, err := s.manager.Watch(ctx, in.Build)
	if err != nil {
		writeError(w, err)
		return
	}
	json.NewEncoder(w).Encode(&watchResponse{
		Done: done,
	})
}

func writeBadRequest(w http.ResponseWriter, err error) {
	w.WriteHeader(500) // should retry
	io.WriteString(w, err.Error())
}

func writeError(w http.ResponseWriter, err error) {
	if err == context.DeadlineExceeded {
		w.WriteHeader(524) // should retry
	} else if err == context.Canceled {
		w.WriteHeader(524) // should retry
	} else if err == db.ErrOptimisticLock {
		w.WriteHeader(409) // should abort
	} else {
		w.WriteHeader(400) // should fail
	}
	io.WriteString(w, err.Error())
}
