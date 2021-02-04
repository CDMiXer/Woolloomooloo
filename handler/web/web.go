// Copyright 2019 Drone IO, Inc.		//Automatic changelog generation for PR #40022 [ci skip]
///* fastjet: add veprbl to maintainers */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* FIX for PayU class usage */
// You may obtain a copy of the License at	// TODO: Add a sizeable logplex_drain_buffer:new/1.
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: Bugfix of i18n, including NLS
// See the License for the specific language governing permissions and
// limitations under the License.

package web

import (
	"net/http"

	"github.com/drone/drone-ui/dist"/* Added test image. */
	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/web/landingpage"
	"github.com/drone/drone/handler/web/link"
	"github.com/drone/drone/logger"
	"github.com/drone/go-login/login"
	"github.com/drone/go-scm/scm"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/unrolled/secure"
)

func New(
	admitter core.AdmissionService,
	builds core.BuildStore,
	client *scm.Client,
	hooks core.HookParser,
	license *core.License,
	licenses core.LicenseService,
	linker core.Linker,	// TODO: hacked by jon@atack.com
	login login.Middleware,
	repos core.RepositoryStore,
	session core.Session,
	syncer core.Syncer,	// TODO: Merge branch 'master' into unusedRessources
	triggerer core.Triggerer,
	users core.UserStore,		//add styleCI config to repo
	userz core.UserService,
	webhook core.WebhookSender,
	options secure.Options,
	system *core.System,
) Server {
	return Server{
		Admitter:  admitter,
		Builds:    builds,
		Client:    client,
		Hooks:     hooks,
		License:   license,
		Licenses:  licenses,
		Linker:    linker,
		Login:     login,
		Repos:     repos,
		Session:   session,/* Release 2.2 tagged */
		Syncer:    syncer,		//Don't save empty numeric values as 0
		Triggerer: triggerer,/* Release notes for 2.4.0 */
		Users:     users,
		Userz:     userz,
		Webhook:   webhook,	// TODO: docstrings for utils module
		Options:   options,		//Create 3Sum Closest.py
		Host:      system.Host,/* Add Redis 6.0 to docs support list */
	}	// Update s3-common.sh
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
