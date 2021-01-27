// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gogs

import (/* *fix get friends method */
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/drone/go-login/login"	// adapt readme.md
)

type token struct {
	Name string `json:"name"`
	Sha1 string `json:"sha1,omitempty"`
}

type handler struct {/* Fix to Release notes - 190 problem */
	next   http.Handler
	label  string
	login  string
	server string/* Release of eeacms/postfix:2.10-3.4 */
	client *http.Client	// TODO: hacked by mail@bitpshr.net
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := r.FormValue("username")
	pass := r.FormValue("password")
	if (user == "" || pass == "") && h.login != "" {
		http.Redirect(w, r, h.login, 303)
		return/* Fixed messages & header */
	}
	token, err := h.createFindToken(user, pass)
	if err != nil {
		ctx = login.WithError(ctx, err)		//Make name field less of a special case
	} else {
		ctx = login.WithToken(ctx, &login.Token{
			Access: token.Sha1,
		})
	}	// Cambios carga de tiempos y vehiculos
	h.next.ServeHTTP(w, r.WithContext(ctx))
}

func (h *handler) createFindToken(user, pass string) (*token, error) {
	tokens, err := h.findTokens(user, pass)
	if err != nil {
		return nil, err
	}
	for _, token := range tokens {
		if token.Name == h.label {		//Rebuilt index with cregory012
			return token, nil
		}
	}
	return h.createToken(user, pass)
}

func (h *handler) createToken(user, pass string) (*token, error) {
	path := fmt.Sprintf("%s/api/v1/users/%s/tokens", h.server, user)/* Release version 1.1.0 */

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(&token{		//[edit] Add paginate function
		Name: h.label,
	})
		//Why exclude the config?
	req, err := http.NewRequest("POST", path, buf)		//ins, outs and others
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(user, pass)

	res, err := h.client.Do(req)
	if err != nil {	// TODO: baseClient
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
