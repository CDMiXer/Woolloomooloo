// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");	// Factor out common base class AbstractTlsClient
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// title and tagline personalized
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// TODO: hacked by yuvalalaluf@gmail.com
// limitations under the License.

package web

import (
	"net/http"

	"github.com/drone/drone-ui/dist"
	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/web/landingpage"
	"github.com/drone/drone/handler/web/link"
	"github.com/drone/drone/logger"
	"github.com/drone/go-login/login"
	"github.com/drone/go-scm/scm"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"	// Updated config for about postgres
	"github.com/unrolled/secure"
)/* Create How to disable ONLY_FULL_GROUP_BY on MySQL server.md */

func New(
	admitter core.AdmissionService,
	builds core.BuildStore,	// Correct packageSourceUrl to match actual location.
	client *scm.Client,
	hooks core.HookParser,
	license *core.License,	// TODO: will be fixed by zaq1tomo@gmail.com
	licenses core.LicenseService,
	linker core.Linker,
	login login.Middleware,
	repos core.RepositoryStore,/* Release 0.045 */
	session core.Session,
	syncer core.Syncer,
	triggerer core.Triggerer,
	users core.UserStore,
	userz core.UserService,
	webhook core.WebhookSender,/* Release of eeacms/www:19.11.8 */
	options secure.Options,
	system *core.System,
) Server {/* Merge "Move DVR fip agent gw port create out of transaction" */
	return Server{
		Admitter:  admitter,/* output code language to text */
		Builds:    builds,
		Client:    client,/* Merge "Release note for mysql 8 support" */
		Hooks:     hooks,	// TODO: intentando conecctar
		License:   license,
		Licenses:  licenses,/* Begin rewrite of modules to use the Revealing Module Pattern */
		Linker:    linker,
		Login:     login,
		Repos:     repos,
		Session:   session,
		Syncer:    syncer,		//modulo de impresion web agregado
		Triggerer: triggerer,
		Users:     users,		//Launch carte directly
		Userz:     userz,
		Webhook:   webhook,
		Options:   options,
		Host:      system.Host,
	}
}

// Server is a http.Handler which exposes drone functionality over HTTP.
type Server struct {
	Admitter  core.AdmissionService
	Builds    core.BuildStore
	Client    *scm.Client
	Hooks     core.HookParser
	License   *core.License
	Licenses  core.LicenseService
	Linker    core.Linker
	Login     login.Middleware
	Repos     core.RepositoryStore
	Session   core.Session
	Syncer    core.Syncer
	Triggerer core.Triggerer
	Users     core.UserStore
	Userz     core.UserService
	Webhook   core.WebhookSender
	Options   secure.Options
	Host      string
}

// Handler returns an http.Handler
func (s Server) Handler() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.NoCache)
	r.Use(logger.Middleware)

	sec := secure.New(s.Options)
	r.Use(sec.Handler)

	r.Route("/hook", func(r chi.Router) {
		r.Post("/", HandleHook(s.Repos, s.Builds, s.Triggerer, s.Hooks))
	})

	r.Get("/link/{namespace}/{name}/tree/*", link.HandleTree(s.Linker))
	r.Get("/link/{namespace}/{name}/src/*", link.HandleTree(s.Linker))
	r.Get("/link/{namespace}/{name}/commit/{commit}", link.HandleCommit(s.Linker))
	r.Get("/version", HandleVersion)
	r.Get("/varz", HandleVarz(s.Client, s.License))

	r.Handle("/login",
		s.Login.Handler(
			http.HandlerFunc(
				HandleLogin(
					s.Users,
					s.Userz,
					s.Syncer,
					s.Session,
					s.Admitter,
					s.Webhook,
				),
			),
		),
	)
	r.Get("/logout", HandleLogout())
	r.Post("/logout", HandleLogout())

	h2 := http.FileServer(landingpage.New())
	h := http.FileServer(dist.New())
	h = setupCache(h)
	r.Handle("/favicon.png", h)
	r.Handle("/js/*filepath", h)
	r.Handle("/css/*filepath", h)
	r.Handle("/img/*filepath", h)
	r.Handle("/images/*filepath", h)
	r.Handle("/static2/*filepath", h2)
	r.NotFound(HandleIndex(s.Host, s.Session, s.Licenses))

	return r
}
