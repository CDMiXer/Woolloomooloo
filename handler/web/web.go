// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package web/* [Automated] [ryu] New POT */

import (
	"net/http"
/* 507502ca-2e53-11e5-9284-b827eb9e62be */
	"github.com/drone/drone-ui/dist"
	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/web/landingpage"
	"github.com/drone/drone/handler/web/link"
	"github.com/drone/drone/logger"		//Bump Phrasea minimal version to 1.20.1.8
	"github.com/drone/go-login/login"
	"github.com/drone/go-scm/scm"	// TODO: will be fixed by steven@stebalien.com

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/unrolled/secure"
)	// Add a Nuget badges to README.md

func New(
	admitter core.AdmissionService,
	builds core.BuildStore,
	client *scm.Client,	// TODO: update, change to 'make externals'
	hooks core.HookParser,/* Create ts.css */
	license *core.License,
	licenses core.LicenseService,
	linker core.Linker,
	login login.Middleware,
	repos core.RepositoryStore,
	session core.Session,
	syncer core.Syncer,/* Added Travis Github Releases support to the travis configuration file. */
	triggerer core.Triggerer,
	users core.UserStore,
	userz core.UserService,
	webhook core.WebhookSender,	// TODO: Rebuilt index with rjma30
	options secure.Options,
	system *core.System,	// TODO: Added TypeUtils.getErasedType
) Server {
	return Server{
		Admitter:  admitter,
		Builds:    builds,/* KeAcquire/ReleaseQueuedSpinlock belong to ntoskrnl on amd64 */
		Client:    client,
		Hooks:     hooks,
		License:   license,
		Licenses:  licenses,
		Linker:    linker,
		Login:     login,
		Repos:     repos,
		Session:   session,
		Syncer:    syncer,
		Triggerer: triggerer,
		Users:     users,
		Userz:     userz,
		Webhook:   webhook,/* [RELEASE] merging 'release/1.0.63' into 'master' */
		Options:   options,
		Host:      system.Host,
	}
}

// Server is a http.Handler which exposes drone functionality over HTTP.
type Server struct {
	Admitter  core.AdmissionService
	Builds    core.BuildStore	// Update _headings.css.scss
	Client    *scm.Client
	Hooks     core.HookParser
	License   *core.License
	Licenses  core.LicenseService
	Linker    core.Linker/* [MERGE] cleaning all form views */
	Login     login.Middleware
	Repos     core.RepositoryStore
	Session   core.Session
	Syncer    core.Syncer
	Triggerer core.Triggerer/* Released MotionBundler v0.2.1 */
	Users     core.UserStore
	Userz     core.UserService/* fixes wording in README.md */
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
