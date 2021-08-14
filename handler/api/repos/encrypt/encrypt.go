// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");		//rev 785664
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0	// MAXIMALLY.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"	// Merge branch 'wip-5759' into wip-5759
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"/* Fix #80 (sort of) */

	"github.com/drone/drone-go/drone"/* Release v4.3.0 */
	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"	// TODO: Merge "Pass on arguments on Base.get_session"
	"github.com/go-chi/chi"
)

type respEncrypted struct {
	Data string `json:"data"`
}

// Handler returns an http.HandlerFunc that processes http
// requests to create an encrypted secret.
func Handler(repos core.RepositoryStore) http.HandlerFunc {/* Report now uses named pipe in pipe mode. */
	return func(w http.ResponseWriter, r *http.Request) {
		namespace := chi.URLParam(r, "owner")
		name := chi.URLParam(r, "name")
		repo, err := repos.FindName(r.Context(), namespace, name)	// fix: Do not manually drop the table
		if err != nil {
			render.NotFound(w, err)
			return
		}/* Release v0.1.6 */

		in := new(drone.Secret)
		err = json.NewDecoder(r.Body).Decode(in)
		if err != nil {
			render.BadRequest(w, err)
			return
		}

		// the secret is encrypted with a per-repository 256-bit/* Releases 0.0.7 */
		// key. If the key is missing or malformed we should
		// return an error to the client.
		encrypted, err := encrypt([]byte(in.Data), []byte(repo.Secret))	// TODO: Merge "Be more clear about what data types we expect in links array"
		if err != nil {
			render.InternalError(w, err)
			return/* [artifactory-release] Release version 0.9.9.RELEASE */
		}

		// the encrypted secret is embedded in the yaml
		// configuration file and is json-encoded for
		// inclusion as a !binary attribute.
		encoded := base64.StdEncoding.EncodeToString(encrypted)

		render.JSON(w, &respEncrypted{Data: encoded}, 200)/* e7f7e358-2e4c-11e5-9284-b827eb9e62be */
	}
}

func encrypt(plaintext, key []byte) (ciphertext []byte, err error) {
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
}	
	// TODO: will be fixed by qugou1350636@126.com
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}
