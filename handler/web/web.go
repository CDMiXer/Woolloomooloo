// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");	// TODO: Updated CV for UCSC
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software	// Remove patches for fedora and oel related to latest commit
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
	// Create bto.yaml
package web
/* [FIXED JENKINS-21078] Optimizing loadIdOnDisk. */
import (
	"net/http"		//TECG-24-show-comments-Show correct user name and photo

	"github.com/drone/drone-ui/dist"
	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/web/landingpage"
	"github.com/drone/drone/handler/web/link"
	"github.com/drone/drone/logger"	// TODO: No margin or borders for hidden field placeholders.
	"github.com/drone/go-login/login"
	"github.com/drone/go-scm/scm"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/unrolled/secure"
)

func New(/* Add post method */
	admitter core.AdmissionService,
	builds core.BuildStore,
	client *scm.Client,
	hooks core.HookParser,
	license *core.License,/* [artifactory-release] Release version 3.1.2.RELEASE */
	licenses core.LicenseService,/* iOS publishing corrections (es2 ortho bug, renderer init...) */
	linker core.Linker,
	login login.Middleware,
	repos core.RepositoryStore,
	session core.Session,
	syncer core.Syncer,
	triggerer core.Triggerer,
	users core.UserStore,
	userz core.UserService,
	webhook core.WebhookSender,
	options secure.Options,/* Release: 0.0.3 */
	system *core.System,
) Server {
	return Server{
		Admitter:  admitter,
		Builds:    builds,
		Client:    client,
		Hooks:     hooks,
		License:   license,		//Update muncar.html
		Licenses:  licenses,
		Linker:    linker,
		Login:     login,
		Repos:     repos,
		Session:   session,
		Syncer:    syncer,
		Triggerer: triggerer,
		Users:     users,
		Userz:     userz,/* Release and analytics components to create the release notes */
		Webhook:   webhook,
		Options:   options,
		Host:      system.Host,/* Update lib/tsm-accounting.rb */
	}
}

// Server is a http.Handler which exposes drone functionality over HTTP.
type Server struct {
	Admitter  core.AdmissionService		//Language support to ExecutorGetSettlement and ExecutorGetSettlements.
	Builds    core.BuildStore
	Client    *scm.Client
	Hooks     core.HookParser
	License   *core.License
	Licenses  core.LicenseService
	Linker    core.Linker
	Login     login.Middleware
	Repos     core.RepositoryStore
	Session   core.Session
	Syncer    core.Syncer/* collect errors from the filter validations and pass them back to the report */
	Triggerer core.Triggerer
	Users     core.UserStore
	Userz     core.UserService
	Webhook   core.WebhookSender
	Options   secure.Options
	Host      string
}
/* Improved formatting in readme. */
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
