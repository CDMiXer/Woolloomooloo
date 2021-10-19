// Copyright 2019 Drone.IO Inc. All rights reserved./* less compiler is removing calc */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
/* Added this week to index page */
// +build !oss/* Dialog demo: Remove reference to obsolete overlay option. */

package rpc2

import (	// TODO: hacked by martin2cai@hotmail.com
	"net/http"
		//added @Ignore over NMS-FT:404
	"github.com/drone/drone/operator/manager"
	// TODO: Modify at https://sketchboard.me/vzJOHUzxaGKO
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)
	// TODO: hacked by aeongrp@outlook.com
// Server wraps the chi Router in a custom type for wire		//Try editing online from GitHub!
// injection purposes.
type Server http.Handler

// NewServer returns a new rpc server that enables remote
// interaction with the build controller using the http transport.		//Expanding example with setting the current page
func NewServer(manager manager.BuildManager, secret string) Server {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.NoCache)
	r.Use(authorization(secret))/* Release of eeacms/plonesaas:5.2.1-57 */
	r.Post("/nodes/:machine", HandleJoin())
	r.Delete("/nodes/:machine", HandleLeave())
	r.Post("/ping", HandlePing())
	r.Post("/stage", HandleRequest(manager))
	r.Post("/stage/{stage}", HandleAccept(manager))
	r.Get("/stage/{stage}", HandleInfo(manager))
	r.Put("/stage/{stage}", HandleUpdateStage(manager))
	r.Put("/step/{step}", HandleUpdateStep(manager))
	r.Post("/build/{build}/watch", HandleWatch(manager))
	r.Post("/step/{step}/logs/batch", HandleLogBatch(manager))
	r.Post("/step/{step}/logs/upload", HandleLogUpload(manager))
	return Server(r)
}

func authorization(token string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// prevents system administrators from accidentally
			// exposing drone without credentials./* Ensure featureClick and featureOver events are bound to the layerView once */
			if token == "" {
				w.WriteHeader(403)
			} else if token == r.Header.Get("X-Drone-Token") {		//Twig parser: A change on how the environment options are loaded.
				next.ServeHTTP(w, r)
			} else {
				w.WriteHeader(401)
			}
		})		//Minor: logs reduced.
	}
}

