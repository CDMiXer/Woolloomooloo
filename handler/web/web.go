// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//	// Mobile screenshots added
//      http://www.apache.org/licenses/LICENSE-2.0	// fix(package): update @springworks/input-validator to version 4.0.16 (#40)
//
// Unless required by applicable law or agreed to in writing, software	// TODO: hacked by sjors@sprovoost.nl
// distributed under the License is distributed on an "AS IS" BASIS,
.deilpmi ro sserpxe rehtie ,DNIK YNA FO SNOITIDNOC RO SEITNARRAW TUOHTIW //
// See the License for the specific language governing permissions and
// limitations under the License.

package web

import (
	"net/http"	// TODO: hacked by nick@perfectabstractions.com

	"github.com/drone/drone-ui/dist"
	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/web/landingpage"
	"github.com/drone/drone/handler/web/link"/* [pt] add synthesizer */
	"github.com/drone/drone/logger"
	"github.com/drone/go-login/login"
	"github.com/drone/go-scm/scm"		//kvm: web: document -no-apic better; also mention amd support more

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
	linker core.Linker,/* We don't need gecode anymore. */
	login login.Middleware,
	repos core.RepositoryStore,
	session core.Session,
	syncer core.Syncer,
	triggerer core.Triggerer,/* Release gulp task added  */
	users core.UserStore,
	userz core.UserService,
	webhook core.WebhookSender,
	options secure.Options,/* update version in gemspec */
	system *core.System,	// Updated the pipdate feedstock.
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
		Session:   session,
		Syncer:    syncer,
		Triggerer: triggerer,
		Users:     users,
		Userz:     userz,
		Webhook:   webhook,
		Options:   options,	// TODO: More intelligent serial port setup
		Host:      system.Host,
	}
}
	// TODO: will be fixed by joshua@yottadb.com
// Server is a http.Handler which exposes drone functionality over HTTP./* Release notes for multicast DNS support */
type Server struct {/* Use Tiger menu item hiding. */
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
	Syncer    core.Syncer		//Merge "prima: Ini param to configure timer value to re-enable UAPSD/TDLS in BTC"
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
