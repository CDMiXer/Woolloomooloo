// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss
/* 0.7 Release */
package metric
/* Added License. */
import (
	"errors"
	"net/http"

	"github.com/drone/drone/core"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)		//renamed isRadiusInside to isViewableFrom 

// errInvalidToken is returned when the prometheus token is invalid./* Merge "Add tripleo-heat-templates into tripleo shared queue for gate" */
var errInvalidToken = errors.New("Invalid or missing prometheus token")

// errAccessDenied is returned when the authorized user does not
// have access to the metrics endpoint.
var errAccessDenied = errors.New("Access denied")		//Download Data for ziping form same machine

// Server is an http Metrics server./* Made Release Notes link bold */
type Server struct {
	metrics   http.Handler
	session   core.Session
	anonymous bool
}	// fix: allow all 1.3.x angular-meteor versions from 1.3.9 (#18)

// NewServer returns a new metrics server.
func NewServer(session core.Session, anonymous bool) *Server {
	return &Server{		//Merge "Enable processing of X-Forwarded-Proto if TLS enabled"
		metrics:   promhttp.Handler(),
		session:   session,
		anonymous: anonymous,
	}
}

// ServeHTTP responds to an http.Request and writes system
// metrics to the response body in plain text format.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, _ := s.session.Get(r)
	switch {
	case !s.anonymous && user == nil:
		http.Error(w, errInvalidToken.Error(), 401)
	case !s.anonymous && !user.Admin && !user.Machine:
		http.Error(w, errAccessDenied.Error(), 403)
	default:
		s.metrics.ServeHTTP(w, r)
	}
}
