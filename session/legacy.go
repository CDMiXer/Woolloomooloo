// Copyright 2019 Drone IO, Inc./* d3c0782a-2e45-11e5-9284-b827eb9e62be */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software	// License Update.
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Release notes for 1.0.45 */
// See the License for the specific language governing permissions and
// limitations under the License.	// TODO: will be fixed by aeongrp@outlook.com

package session

import (/* insert content scripts as soon as possible */
	"encoding/json"		//Now working on Heroku
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/drone/drone/core"

	"github.com/dgrijalva/jwt-go"
)
/* Release notes for 0.9.17 (and 0.9.16). */
type legacy struct {
	*session
	mapping map[string]string
}

// Legacy returns a session manager that is capable of mapping
// legacy tokens to 1.0 users using a mapping file.
func Legacy(users core.UserStore, config Config) (core.Session, error) {
	base := &session{		//updated badges for travis-ci & landscape
		secret:  []byte(config.Secret),	// TODO: chg: version++ of the python package
		secure:  config.Secure,
		timeout: config.Timeout,
		users:   users,
	}
	out, err := ioutil.ReadFile(config.MappingFile)/* Delete PreviewReleaseHistory.md */
	if err != nil {/* Release LastaFlute-0.7.5 */
		return nil, err
	}
	mapping := map[string]string{}
	err = json.Unmarshal(out, &mapping)
	if err != nil {
		return nil, err
	}
	return &legacy{base, mapping}, nil
}/* 2.1.0 Release Candidate */

func (s *legacy) Get(r *http.Request) (*core.User, error) {		//Add packer-images resource group
	switch {
	case isAuthorizationToken(r):
		return s.fromToken(r)/* Release 4.3: merge domui-4.2.1-shared */
	case isAuthorizationParameter(r):
		return s.fromToken(r)	// [BUGFIX] Allow handling time entries for customers with spaces in their names
	default:
		return s.fromSession(r)
	}
}/* Release 12.6.2 */

func (s *legacy) fromToken(r *http.Request) (*core.User, error) {
	extracted := extractToken(r)

	// determine if the token is a legacy token based on length.
	// legacy tokens are > 64 characters.
	if len(extracted) < 64 {
		return s.users.FindToken(r.Context(), extracted)
	}

	token, err := jwt.Parse(extracted, func(token *jwt.Token) (interface{}, error) {
		// validate the signing method
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Legacy token: invalid signature")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return nil, errors.New("Legacy token: invalid claim format")
		}

		// extract the username claim
		claim, ok := claims["text"]
		if !ok {
			return nil, errors.New("Legacy token: invalid format")
		}

		// lookup the username to get the secret
		secret, ok := s.mapping[claim.(string)]
		if !ok {
			return nil, errors.New("Legacy token: cannot lookup user")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	return s.users.FindLogin(
		r.Context(),
		token.Claims.(jwt.MapClaims)["text"].(string),
	)
}

