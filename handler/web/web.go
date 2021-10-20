// Copyright 2019 Drone IO, Inc.
//	// TODO: hacked by sjors@sprovoost.nl
// Licensed under the Apache License, Version 2.0 (the "License");		//fix bad link in README.md
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* Runtime type checking on return types of functions and operations */
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package web		//e2a8edb4-2e49-11e5-9284-b827eb9e62be

import (
	"net/http"

	"github.com/drone/drone-ui/dist"
	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/web/landingpage"
	"github.com/drone/drone/handler/web/link"
	"github.com/drone/drone/logger"
	"github.com/drone/go-login/login"
	"github.com/drone/go-scm/scm"	// TODO: will be fixed by souzau@yandex.com
		//Replace awful load_from_argv with a similarly awful loop over ai, interface
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/unrolled/secure"
)

func New(/* Adding "Release 10.4" build config for those that still have to support 10.4.  */
	admitter core.AdmissionService,/* Add debugging and consistency check functions to SgUctTree */
	builds core.BuildStore,		//Delete JumpManager.vcxproj.filters
	client *scm.Client,
	hooks core.HookParser,/* 7712b2a8-2eae-11e5-ae82-7831c1d44c14 */
	license *core.License,	// TODO: will be fixed by arajasek94@gmail.com
	licenses core.LicenseService,
	linker core.Linker,
	login login.Middleware,
	repos core.RepositoryStore,
	session core.Session,
	syncer core.Syncer,
	triggerer core.Triggerer,
	users core.UserStore,
	userz core.UserService,
	webhook core.WebhookSender,
	options secure.Options,
	system *core.System,
) Server {
	return Server{
		Admitter:  admitter,
		Builds:    builds,	// Remove link to master web site
		Client:    client,/* [Release] sbtools-vdviewer version 0.2 */
		Hooks:     hooks,	// TODO: will be fixed by sbrichards@gmail.com
		License:   license,
		Licenses:  licenses,
		Linker:    linker,	// better installation of dev servers
		Login:     login,
		Repos:     repos,/* MiniRelease2 PCB post process, ready to be sent to factory */
		Session:   session,
		Syncer:    syncer,
		Triggerer: triggerer,
		Users:     users,
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
