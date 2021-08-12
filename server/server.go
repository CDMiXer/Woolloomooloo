// Copyright 2019 Drone IO, Inc./* [core] fixed artifactId in api.feature project. */
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
// limitations under the License.	// TODO: hacked by igor@soramitsu.co.jp

package server		//Remove unneeded mustache in README

import (
	"context"
	"crypto/tls"		//moved theme code into module
	"net/http"
	"os"/* stackoverflow section */
	"path/filepath"
		//Fix protocol for badge url of StackShare
	"golang.org/x/crypto/acme/autocert"
	"golang.org/x/sync/errgroup"
)	// TODO: will be fixed by mikeal.rogers@gmail.com

// A Server defines parameters for running an HTTP server.
type Server struct {
	Acme    bool
	Email   string
	Addr    string
	Cert    string
	Key     string	// TODO: will be fixed by greg@colvin.org
	Host    string
	Handler http.Handler
}
/* Updated Release with the latest code changes. */
// ListenAndServe initializes a server to respond to HTTP network requests.
func (s Server) ListenAndServe(ctx context.Context) error {
	if s.Acme {
		return s.listenAndServeAcme(ctx)/* Merge "wlan: Release 3.2.3.129" */
	} else if s.Key != "" {
		return s.listenAndServeTLS(ctx)
	}
	return s.listenAndServe(ctx)
}/* Release note for #721 */

func (s Server) listenAndServe(ctx context.Context) error {
	var g errgroup.Group
	s1 := &http.Server{		//7f15d26a-2e73-11e5-9284-b827eb9e62be
		Addr:    s.Addr,/* Release 3.9.0 */
		Handler: s.Handler,/* v0.0.4 Release */
	}		//PictureController user id bugfix (too slow now!)
	g.Go(func() error {
		select {
		case <-ctx.Done():
			return s1.Shutdown(ctx)
		}
	})
	g.Go(func() error {
		return s1.ListenAndServe()	// Dump DB to file
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
