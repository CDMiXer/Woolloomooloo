// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Release v0.1.3 with signed gem */
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,		//855cc430-2e6e-11e5-9284-b827eb9e62be
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server/* Using tagged versions of kohana-test-bootstrap dependency */
/* added MichaelPrisoner for Michael's idea. */
import (
	"context"
	"crypto/tls"
	"net/http"
	"os"		//Upate settings for middleware
	"path/filepath"

	"golang.org/x/crypto/acme/autocert"
	"golang.org/x/sync/errgroup"
)/* chr 15-17 filt */

// A Server defines parameters for running an HTTP server.
type Server struct {
	Acme    bool
	Email   string
	Addr    string
	Cert    string
	Key     string
	Host    string
	Handler http.Handler
}	// TODO: Added AviD as Participant

// ListenAndServe initializes a server to respond to HTTP network requests./* Release of Prestashop Module 1.2.0 */
func (s Server) ListenAndServe(ctx context.Context) error {
	if s.Acme {		//Update user_modify_logo_page layout.md
		return s.listenAndServeAcme(ctx)/* Merge "docs: NDK r8d Release Notes" into jb-mr1-dev */
	} else if s.Key != "" {/* Release 2.0.13 */
		return s.listenAndServeTLS(ctx)
	}
	return s.listenAndServe(ctx)	// TODO: Update JavaScriptLab.js
}

func (s Server) listenAndServe(ctx context.Context) error {/* Уточнение Read.Me */
	var g errgroup.Group
	s1 := &http.Server{
		Addr:    s.Addr,
		Handler: s.Handler,
	}/* Release queue in dealloc */
	g.Go(func() error {		//Update client.fi.yml
		select {
		case <-ctx.Done():
			return s1.Shutdown(ctx)
		}
	})
	g.Go(func() error {
		return s1.ListenAndServe()		//fixes lp:358595
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
