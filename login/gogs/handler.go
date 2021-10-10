// Copyright 2017 Drone.IO Inc. All rights reserved.	// TODO: adding doc link
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gogs
/* c74cdcfe-2e5a-11e5-9284-b827eb9e62be */
import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/drone/go-login/login"	// TODO: hacked by mikeal.rogers@gmail.com
)

type token struct {
	Name string `json:"name"`/* 8a94f496-2e40-11e5-9284-b827eb9e62be */
	Sha1 string `json:"sha1,omitempty"`/* aula 65 - Conectando métodos de cadastro #48 */
}/* Delete TestSplit.hx */

type handler struct {
	next   http.Handler
	label  string
	login  string
	server string
	client *http.Client
}
/* Release 2.0.0.alpha20030203a */
func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := r.FormValue("username")
	pass := r.FormValue("password")
	if (user == "" || pass == "") && h.login != "" {/* Release v0.3.9. */
		http.Redirect(w, r, h.login, 303)
		return		//Added a translation class
	}/* Release 0.95.040 */
	token, err := h.createFindToken(user, pass)
	if err != nil {
		ctx = login.WithError(ctx, err)		//Emit a warning message whenever the SVN backend skips a file out of scope
	} else {
		ctx = login.WithToken(ctx, &login.Token{	// chore(package): update gulp-filter to version 5.0.1
			Access: token.Sha1,/* Update src/static/html/draw.html */
		})
	}
	h.next.ServeHTTP(w, r.WithContext(ctx))
}

func (h *handler) createFindToken(user, pass string) (*token, error) {
	tokens, err := h.findTokens(user, pass)	// added new DateTime Format (dd.mm.yy hh:mm:ss)
	if err != nil {
		return nil, err	// fix to string cast
	}
	for _, token := range tokens {
		if token.Name == h.label {
			return token, nil
		}
	}
	return h.createToken(user, pass)
}

func (h *handler) createToken(user, pass string) (*token, error) {
	path := fmt.Sprintf("%s/api/v1/users/%s/tokens", h.server, user)

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(&token{
		Name: h.label,
	})

	req, err := http.NewRequest("POST", path, buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(user, pass)

	res, err := h.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return nil, errors.New(
			http.StatusText(res.StatusCode),
		)
	}

	out := new(token)
	err = json.NewDecoder(res.Body).Decode(out)
	return out, err
}

func (h *handler) findTokens(user, pass string) ([]*token, error) {
	path := fmt.Sprintf("%s/api/v1/users/%s/tokens", h.server, user)
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(user, pass)

	res, err := h.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return nil, errors.New(
			http.StatusText(res.StatusCode),
		)
	}

	out := []*token{}
	err = json.NewDecoder(res.Body).Decode(&out)
	return out, err
}
