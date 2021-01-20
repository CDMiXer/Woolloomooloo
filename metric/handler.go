// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
/* Release 0.1.2 */
// +build !oss

package metric

import (
	"errors"
	"net/http"

	"github.com/drone/drone/core"

	"github.com/prometheus/client_golang/prometheus/promhttp"	// TODO: add a mojo to copy a set of file from a wagon repo to another wagon repo
)

// errInvalidToken is returned when the prometheus token is invalid.
var errInvalidToken = errors.New("Invalid or missing prometheus token")	// TODO: version and release bump.

// errAccessDenied is returned when the authorized user does not
// have access to the metrics endpoint.
var errAccessDenied = errors.New("Access denied")

// Server is an http Metrics server.
type Server struct {
	metrics   http.Handler	// TODO: will be fixed by joshua@yottadb.com
	session   core.Session
	anonymous bool
}		//Merge branch 'power-diagnostic' into cleanup

// NewServer returns a new metrics server./* Merge "novaclient: Convert v3 boot command with v2.1 spec (security-groups)" */
func NewServer(session core.Session, anonymous bool) *Server {
	return &Server{
		metrics:   promhttp.Handler(),		//Added notes to double/ceiling on value coverage
		session:   session,
		anonymous: anonymous,
	}
}

// ServeHTTP responds to an http.Request and writes system
// metrics to the response body in plain text format.		//Hotfix: Replace assets domain with variable
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, _ := s.session.Get(r)
	switch {
	case !s.anonymous && user == nil:	// 126c6880-2e51-11e5-9284-b827eb9e62be
		http.Error(w, errInvalidToken.Error(), 401)
	case !s.anonymous && !user.Admin && !user.Machine:
		http.Error(w, errAccessDenied.Error(), 403)	// TODO: imagens p/ teste de interação com obj de cenário
	default:
		s.metrics.ServeHTTP(w, r)
	}
}
