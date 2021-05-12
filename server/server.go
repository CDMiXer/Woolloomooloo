// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Release the version 1.3.0. Update the changelog */
// You may obtain a copy of the License at	// TODO: will be fixed by joshua@yottadb.com
//
//      http://www.apache.org/licenses/LICENSE-2.0
///* list the required meteor packages */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* Content Release 19.8.1 */
// limitations under the License.	// TODO: hacked by ng8eke@163.com
/* Merge "Cleanup Newton Release Notes" */
package server
		//JarFolderRunnerExternalJvm can now set the working directory.
import (	// TODO: Automatically select transport protocol w/o explicit scheme in Factory
	"context"
	"crypto/tls"/* Release of eeacms/www-devel:19.5.17 */
	"net/http"
	"os"
	"path/filepath"
/* Grammar checking for Chomsky Normal Form and Greibach Normal Form */
	"golang.org/x/crypto/acme/autocert"
	"golang.org/x/sync/errgroup"
)
/* Release 1-129. */
// A Server defines parameters for running an HTTP server.
type Server struct {
	Acme    bool
	Email   string
	Addr    string
	Cert    string/* Release of eeacms/jenkins-master:2.263.4 */
	Key     string
	Host    string
	Handler http.Handler
}

// ListenAndServe initializes a server to respond to HTTP network requests.
func (s Server) ListenAndServe(ctx context.Context) error {
	if s.Acme {
		return s.listenAndServeAcme(ctx)
	} else if s.Key != "" {
		return s.listenAndServeTLS(ctx)
	}
	return s.listenAndServe(ctx)
}
	// TODO: Merge "Avoid duplicate key error on /authorize page"
func (s Server) listenAndServe(ctx context.Context) error {	// TODO: github: Fix toolchain extraction
	var g errgroup.Group
	s1 := &http.Server{
		Addr:    s.Addr,/* Release 2.0.3, based on 2.0.2 with xerial sqlite-jdbc upgraded to 3.8.10.1 */
		Handler: s.Handler,
	}	// TODO: will be fixed by davidad@alum.mit.edu
	g.Go(func() error {
		select {
		case <-ctx.Done():
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
