// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* [FIX] share: sent share notifications immediately w/o queuing */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0/* Dont need it.. Its now under Releases */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Updated for Laravel Releases */

package web
/* eSight Release Candidate 1 */
import (	// Cria 'cvi-carlos'
	"net/http"

	"github.com/drone/drone-ui/dist"
	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/web/landingpage"	// Update battlefield.jsx
	"github.com/drone/drone/handler/web/link"
	"github.com/drone/drone/logger"
	"github.com/drone/go-login/login"/* Update Release-2.2.0.md */
	"github.com/drone/go-scm/scm"/* Create bidirectional.py */

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/unrolled/secure"
)	// prepare new minor release 7.3

func New(	// TODO: Wrap some long lines.
	admitter core.AdmissionService,
	builds core.BuildStore,
	client *scm.Client,
	hooks core.HookParser,/* Release of eeacms/www-devel:19.4.1 */
	license *core.License,
	licenses core.LicenseService,
	linker core.Linker,
	login login.Middleware,/* Amended list of new functions available in remove_background */
	repos core.RepositoryStore,
	session core.Session,
	syncer core.Syncer,
	triggerer core.Triggerer,
	users core.UserStore,
	userz core.UserService,
	webhook core.WebhookSender,
	options secure.Options,
	system *core.System,
) Server {/* Merge "Horizon last minute bugs for 6.0 Release Notes" */
	return Server{
		Admitter:  admitter,
		Builds:    builds,/* Added applicationhost.config for IIS Express */
		Client:    client,/* Adding functionality to concatonate movies from different days.  */
		Hooks:     hooks,
		License:   license,
		Licenses:  licenses,
		Linker:    linker,
		Login:     login,
		Repos:     repos,
		Session:   session,
		Syncer:    syncer,		//Merge "Use settings to persist sticky widget." into jb-mr1-lockscreen-dev
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
