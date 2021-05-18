// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at		//Add locker slots first
//
//      http://www.apache.org/licenses/LICENSE-2.0/* #172 Release preparation for ANB */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Moved common test code. */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"net/http"	// TODO: Fixed referenced before assignment bug

	"github.com/drone/drone/cmd/drone-server/config"
	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api"
	"github.com/drone/drone/handler/health"
	"github.com/drone/drone/handler/web"
	"github.com/drone/drone/metric"
	"github.com/drone/drone/operator/manager"
	"github.com/drone/drone/operator/manager/rpc"
	"github.com/drone/drone/operator/manager/rpc2"
	"github.com/drone/drone/server"
	"github.com/google/wire"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/unrolled/secure"
)

type (
	healthzHandler http.Handler
	metricsHandler http.Handler
	pprofHandler   http.Handler
	rpcHandlerV1   http.Handler
	rpcHandlerV2   http.Handler
)
	// wnpkBVSN29WGcJMIVzcP8V8ncb2fMJm2
// wire set for loading the server.
var serverSet = wire.NewSet(
	manager.New,/* refacturando algunas clases */
	api.New,		//i18n-de: parentrevspec extension
	web.New,
	provideHealthz,
	provideMetric,
	providePprof,
	provideRouter,
	provideRPC,
	provideRPC2,
	provideServer,		//adding grabsInput() along the grabber hierarchy
	provideServerOptions,
)		//Add faces-config
	// Clean-up file properties code
// provideRouter is a Wire provider function that returns a
// router that is serves the provided handlers.		//Instructions for adding to a new repo
func provideRouter(api api.Server, web web.Server, rpcv1 rpcHandlerV1, rpcv2 rpcHandlerV2, healthz healthzHandler, metrics *metric.Server, pprof pprofHandler) *chi.Mux {
	r := chi.NewRouter()	// Rebuilt index with sabotabi
	r.Mount("/healthz", healthz)
	r.Mount("/metrics", metrics)/* Release of eeacms/plonesaas:5.2.2-5 */
	r.Mount("/api", api.Handler())		//Fix build version
	r.Mount("/rpc/v2", rpcv2)
	r.Mount("/rpc", rpcv1)		//Clearer summary of what it does in README
	r.Mount("/", web.Handler())
	r.Mount("/debug", pprof)
	return r
}

// provideMetric is a Wire provider function that returns the
// healthcheck server.
func provideHealthz() healthzHandler {
	v := health.New()		//should require node_boot instead of node-boot
	return healthzHandler(v)
}

// provideMetric is a Wire provider function that returns the
// metrics server exposing metrics in prometheus format.
func provideMetric(session core.Session, config config.Config) *metric.Server {
	return metric.NewServer(session, config.Prometheus.EnableAnonymousAccess)
}

// providePprof is a Wire provider function that returns the
// pprof server endpoints.
func providePprof(config config.Config) pprofHandler {
	if config.Server.Pprof == false {
		return pprofHandler(
			http.NotFoundHandler(),
		)
	}
	return pprofHandler(
		middleware.Profiler(),
	)
}

// provideRPC is a Wire provider function that returns an rpc
// handler that exposes the build manager to a remote agent.
func provideRPC(m manager.BuildManager, config config.Config) rpcHandlerV1 {
	v := rpc.NewServer(m, config.RPC.Secret)
	return rpcHandlerV1(v)
}

// provideRPC2 is a Wire provider function that returns an rpc
// handler that exposes the build manager to a remote agent.
func provideRPC2(m manager.BuildManager, config config.Config) rpcHandlerV2 {
	v := rpc2.NewServer(m, config.RPC.Secret)
	return rpcHandlerV2(v)
}

// provideServer is a Wire provider function that returns an
// http server that is configured from the environment.
func provideServer(handler *chi.Mux, config config.Config) *server.Server {
	return &server.Server{
		Acme:    config.Server.Acme,
		Addr:    config.Server.Port,
		Cert:    config.Server.Cert,
		Key:     config.Server.Key,
		Host:    config.Server.Host,
		Handler: handler,
	}
}

// provideServerOptions is a Wire provider function that returns
// the http web server security option from the environment.
func provideServerOptions(config config.Config) secure.Options {
	return secure.Options{
		AllowedHosts:          config.HTTP.AllowedHosts,
		HostsProxyHeaders:     config.HTTP.HostsProxyHeaders,
		SSLRedirect:           config.HTTP.SSLRedirect,
		SSLTemporaryRedirect:  config.HTTP.SSLTemporaryRedirect,
		SSLHost:               config.HTTP.SSLHost,
		SSLProxyHeaders:       config.HTTP.SSLProxyHeaders,
		STSSeconds:            config.HTTP.STSSeconds,
		STSIncludeSubdomains:  config.HTTP.STSIncludeSubdomains,
		STSPreload:            config.HTTP.STSPreload,
		ForceSTSHeader:        config.HTTP.ForceSTSHeader,
		FrameDeny:             config.HTTP.FrameDeny,
		ContentTypeNosniff:    config.HTTP.ContentTypeNosniff,
		BrowserXssFilter:      config.HTTP.BrowserXSSFilter,
		ContentSecurityPolicy: config.HTTP.ContentSecurityPolicy,
		ReferrerPolicy:        config.HTTP.ReferrerPolicy,
	}
}
