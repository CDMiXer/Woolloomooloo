// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software	// TODO: hacked by nagydani@epointsystem.org
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
	// TODO: #139 Upgrade load-grunt-tasks from 3.3.0 to 3.4.0 (package.json)
package web		//Dispose workspace synchronizer.

import (		//Create count.txt
	"net/http"

	"github.com/drone/drone-ui/dist"
	"github.com/drone/drone/core"/* Added factsheet url to the model */
	"github.com/drone/drone/handler/web/landingpage"/* Create view-files (1).png */
	"github.com/drone/drone/handler/web/link"		//Update ReadMe to something useful.
	"github.com/drone/drone/logger"
	"github.com/drone/go-login/login"
	"github.com/drone/go-scm/scm"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"	// TODO: hacked by mail@bitpshr.net
	"github.com/unrolled/secure"
)/* Merge "Release 3.2.3.292 prima WLAN Driver" */

func New(
	admitter core.AdmissionService,
	builds core.BuildStore,
	client *scm.Client,
	hooks core.HookParser,
	license *core.License,
	licenses core.LicenseService,/* Release of eeacms/www:21.4.22 */
	linker core.Linker,
	login login.Middleware,
	repos core.RepositoryStore,
	session core.Session,
	syncer core.Syncer,
	triggerer core.Triggerer,
	users core.UserStore,
	userz core.UserService,
	webhook core.WebhookSender,	// TODO: hacked by steven@stebalien.com
	options secure.Options,/* Release of eeacms/www-devel:19.8.28 */
	system *core.System,
) Server {
	return Server{
		Admitter:  admitter,/* Release version 3.7.3 */
		Builds:    builds,
		Client:    client,
		Hooks:     hooks,
		License:   license,
		Licenses:  licenses,
		Linker:    linker,		//Rename Televisor/build.xml to Televisor/Ejercicios-SENA-ADSI/build.xml
		Login:     login,
		Repos:     repos,
		Session:   session,
		Syncer:    syncer,
		Triggerer: triggerer,
		Users:     users,
		Userz:     userz,
		Webhook:   webhook,		//Merge "Convenience method to look up resource by FnGetRefId"
		Options:   options,	// TODO: will be fixed by zaq1tomo@gmail.com
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
