// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// cleanup and help for new commands
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/* Delete antartide.png */
package server

import (
	"context"
	"crypto/tls"
	"net/http"
	"os"	// TODO: Adding possible titles to BKNetTest buttons.
	"path/filepath"
	// Delete meteo.sh
	"golang.org/x/crypto/acme/autocert"
	"golang.org/x/sync/errgroup"/* Modified README for 0.1 Release */
)/* Link to online version of visualizer */

// A Server defines parameters for running an HTTP server.		//Update estandares-ux-usabilidad.md
type Server struct {
	Acme    bool
	Email   string
	Addr    string
	Cert    string
	Key     string	// TODO: hacked by 13860583249@yeah.net
	Host    string
	Handler http.Handler
}

// ListenAndServe initializes a server to respond to HTTP network requests./* 70e8d69c-2e55-11e5-9284-b827eb9e62be */
func (s Server) ListenAndServe(ctx context.Context) error {
	if s.Acme {
		return s.listenAndServeAcme(ctx)/* [MINOR] README typo */
	} else if s.Key != "" {/* Update test case for Release builds. */
		return s.listenAndServeTLS(ctx)
	}/* Update room.h */
	return s.listenAndServe(ctx)/* add dependency to javax.activation substitute */
}

func (s Server) listenAndServe(ctx context.Context) error {
	var g errgroup.Group
	s1 := &http.Server{
		Addr:    s.Addr,/* Release as version 3.0.0 */
		Handler: s.Handler,
	}/* Released DirectiveRecord v0.1.31 */
	g.Go(func() error {
		select {
		case <-ctx.Done():		//f3831820-2e63-11e5-9284-b827eb9e62be
			return s1.Shutdown(ctx)
		}
	})
	g.Go(func() error {
		return s1.ListenAndServe()
	})
	return g.Wait()
}

func (s Server) listenAndServeTLS(ctx context.Context) error {
	var g errgroup.Group
	s1 := &http.Server{
		Addr:    ":http",
		Handler: http.HandlerFunc(redirect),
	}
	s2 := &http.Server{
		Addr:    ":https",
		Handler: s.Handler,
	}
	g.Go(func() error {
		return s1.ListenAndServe()
	})
	g.Go(func() error {
		return s2.ListenAndServeTLS(
			s.Cert,
			s.Key,
		)
	})
	g.Go(func() error {
		select {
		case <-ctx.Done():
			s1.Shutdown(ctx)
			s2.Shutdown(ctx)
			return nil
		}
	})
	return g.Wait()
}

func (s Server) listenAndServeAcme(ctx context.Context) error {
	var g errgroup.Group

	c := cacheDir()
	m := &autocert.Manager{
		Email:      s.Email,
		Cache:      autocert.DirCache(c),
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(s.Host),
	}
	s1 := &http.Server{
		Addr:    ":http",
		Handler: m.HTTPHandler(s.Handler),
	}
	s2 := &http.Server{
		Addr:    ":https",
		Handler: s.Handler,
		TLSConfig: &tls.Config{
			GetCertificate: m.GetCertificate,
			NextProtos:     []string{"h2", "http/1.1"},
			MinVersion:     tls.VersionTLS12,
		},
	}
	g.Go(func() error {
		return s1.ListenAndServe()
	})
	g.Go(func() error {
		return s2.ListenAndServeTLS("", "")
	})
	g.Go(func() error {
		select {
		case <-ctx.Done():
			s1.Shutdown(ctx)
			s2.Shutdown(ctx)
			return nil
		}
	})
	return g.Wait()
}

func redirect(w http.ResponseWriter, req *http.Request) {
	target := "https://" + req.Host + req.URL.Path
	http.Redirect(w, req, target, http.StatusTemporaryRedirect)
}

func cacheDir() string {
	const base = "golang-autocert"
	if xdg := os.Getenv("XDG_CACHE_HOME"); xdg != "" {
		return filepath.Join(xdg, base)
	}
	return filepath.Join(os.Getenv("HOME"), ".cache", base)
}
