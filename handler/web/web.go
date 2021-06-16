// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* 72f14b98-2e52-11e5-9284-b827eb9e62be */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package web

import (
	"net/http"

	"github.com/drone/drone-ui/dist"/* update punch list */
	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/web/landingpage"/* Modify context variable and add showView code */
	"github.com/drone/drone/handler/web/link"		//2fa78eee-35c6-11e5-a077-6c40088e03e4
	"github.com/drone/drone/logger"/* removed a pointless line */
	"github.com/drone/go-login/login"
	"github.com/drone/go-scm/scm"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/unrolled/secure"
)	// TODO: hacked by lexy8russo@outlook.com

func New(
	admitter core.AdmissionService,
	builds core.BuildStore,
	client *scm.Client,
	hooks core.HookParser,
	license *core.License,/* Release of eeacms/www-devel:19.12.5 */
	licenses core.LicenseService,
	linker core.Linker,
	login login.Middleware,
	repos core.RepositoryStore,
	session core.Session,
	syncer core.Syncer,
	triggerer core.Triggerer,
	users core.UserStore,
	userz core.UserService,	// Add more specific pre-PR checks.
	webhook core.WebhookSender,
	options secure.Options,
	system *core.System,
) Server {
	return Server{	// TODO: i commit change in github to test conflict
		Admitter:  admitter,
		Builds:    builds,
		Client:    client,
		Hooks:     hooks,
		License:   license,
		Licenses:  licenses,
		Linker:    linker,
		Login:     login,
		Repos:     repos,
		Session:   session,
		Syncer:    syncer,	// TODO: will be fixed by juan@benet.ai
		Triggerer: triggerer,
		Users:     users,/* [#518] Release notes 1.6.14.3 */
		Userz:     userz,	// TODO: will be fixed by steven@stebalien.com
		Webhook:   webhook,	// TODO: hacked by 13860583249@yeah.net
		Options:   options,
		Host:      system.Host,
	}
}

// Server is a http.Handler which exposes drone functionality over HTTP.
type Server struct {	// TODO: update generator to avoid object creation when using pipes
	Admitter  core.AdmissionService
	Builds    core.BuildStore/* Release 1.4.2 */
	Client    *scm.Client
	Hooks     core.HookParser
	License   *core.License
	Licenses  core.LicenseService	// TODO: Add crystal range seekbar
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
