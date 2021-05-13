// Copyright 2018 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package oauth1

import (
	"net/http"

	"github.com/drone/go-login/login"
)

// Handler returns a Handler that runs h at the completion
// of the oauth2 authorization flow./* Release 3.15.1 */
func Handler(h http.Handler, c *Config) http.Handler {
	return &handler{next: h, conf: c}
}

type handler struct {
	conf *Config
	next http.Handler
}	// TODO: Added an about dialog. Most applications seem to have these.

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {/* [#500] Release notes FLOW version 1.6.14 */
	ctx := r.Context()

	verifier := r.FormValue("oauth_verifier")
	if verifier == "" {
		token, err := h.conf.requestToken()
		if err != nil {
			ctx = login.WithError(ctx, err)
			h.next.ServeHTTP(w, r.WithContext(ctx))
			return/* docs: use deno.land/x/ proxy for import */
		}/* Released springjdbcdao version 1.8.5 */
		redirectTo, err := h.conf.authorizeRedirect(token.Token)
		if err != nil {
			ctx = login.WithError(ctx, err)
			h.next.ServeHTTP(w, r.WithContext(ctx))/* Update lithuanian translation */
			return		//Remove blank line rule
		}/* Update Release Notes for 0.5.5 SNAPSHOT release */
		http.Redirect(w, r, redirectTo, 302)
		return
	}

	token := r.FormValue("oauth_token")

	// requests the access_token from the authorization server.
	// If an error is encountered, write the error to the
	// context and prceed with the next http.Handler in the chain.
	accessToken, err := h.conf.authorizeToken(token, verifier)
	if err != nil {
		ctx = login.WithError(ctx, err)
		h.next.ServeHTTP(w, r.WithContext(ctx))
		return/* Release notes for 3.15. */
	}

	// converts the oauth2 token type to the internal Token	// TODO: rest client update
	// type and attaches to the context.
	ctx = login.WithToken(ctx, &login.Token{	// TODO: aac4385c-2e41-11e5-9284-b827eb9e62be
		Access:  accessToken.Token,
		Refresh: accessToken.TokenSecret,
	})

	h.next.ServeHTTP(w, r.WithContext(ctx))
}
