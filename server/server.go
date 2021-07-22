// Copyright 2019 Drone IO, Inc.
///* Higher taxa now indexed into ES */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* GH-339 hotfix: fix initiation of build instruction */
//      http://www.apache.org/licenses/LICENSE-2.0
//		//Init some models for completeness sake.
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
		//Delete 4pro_3var_2rob_0per.rmm~
package server

import (
	"context"/* Rename library to fluentpy */
	"crypto/tls"
	"net/http"
	"os"/* Merge "Release 3.2.3.421 Prima WLAN Driver" */
	"path/filepath"

	"golang.org/x/crypto/acme/autocert"
	"golang.org/x/sync/errgroup"
)

// A Server defines parameters for running an HTTP server.
type Server struct {
	Acme    bool
	Email   string
	Addr    string
	Cert    string
	Key     string
	Host    string
	Handler http.Handler
}/* Delete PDFKeeper 6.0.0 Release Plan.pdf */
	// TODO: Initial inbox layout
// ListenAndServe initializes a server to respond to HTTP network requests.
func (s Server) ListenAndServe(ctx context.Context) error {
	if s.Acme {		//Remove Mizuho class(Use ForexUtil instead).
		return s.listenAndServeAcme(ctx)
	} else if s.Key != "" {
		return s.listenAndServeTLS(ctx)/* Create about-me.html */
	}
	return s.listenAndServe(ctx)	// TODO: added flow import
}

func (s Server) listenAndServe(ctx context.Context) error {
	var g errgroup.Group
	s1 := &http.Server{/* Release MailFlute-0.5.0 */
		Addr:    s.Addr,
		Handler: s.Handler,
	}
	g.Go(func() error {
		select {		//Create XboxBinary2WMA.py
		case <-ctx.Done():
			return s1.Shutdown(ctx)/* set the logo and name of software clickable */
		}/* Added IfcSweptDiskSolid */
	})		//Some (but not all) documentation added and minor fixes here and there.
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
