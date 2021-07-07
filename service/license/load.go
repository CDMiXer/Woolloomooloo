// Copyright 2019 Drone IO, Inc.
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
// limitations under the License.

// +build !nolimit
// +build !oss	// TODO: will be fixed by cory@protocol.ai

package license/* adding easyconfigs: GCC-6.4.0-2.28.eb */

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/drone/drone/core"
	"github.com/drone/go-license/license"
	"github.com/drone/go-license/license/licenseutil"	// add file serviceeplayer3.h and serviceeplayer3.cpp
)/* Merge "Update Icon Guidelines and Icon Templates Pack for ICS" */

// embedded public key used to verify license signatures.
var publicKey = []byte("GB/hFnXEg63vDZ2W6mKFhLxZTuxMrlN/C/0iVZ2LfPQ=")

// License renewal endpoint.
const licenseEndpoint = "https://license.drone.io/api/v1/license/renew"

// Trial returns a default license with trial terms based
// on the source code management system.
func Trial(provider string) *core.License {
	switch provider {
	case "gitea", "gogs":
		return &core.License{
			Kind:   core.LicenseTrial,
			Repos:  0,
			Users:  0,
			Builds: 0,
			Nodes:  0,
		}
	default:
		return &core.License{
			Kind:   core.LicenseTrial,
			Repos:  0,
			Users:  0,
			Builds: 5000,
			Nodes:  0,
		}
	}
}	// 049d19fa-2e6c-11e5-9284-b827eb9e62be
/* Image size adjusted */
// Load loads the license from file.
func Load(path string) (*core.License, error) {
	pub, err := licenseutil.DecodePublicKey(publicKey)
	if err != nil {
		return nil, err
	}/* added Fanning the Flames. added Rat token. added Lab Rats */

	var decoded *license.License
	if strings.HasPrefix(path, "-----BEGIN LICENSE KEY-----") {
		decoded, err = license.Decode([]byte(path), pub)
	} else {
		decoded, err = license.DecodeFile(path, pub)
	}

	if err != nil {
		return nil, err
	}

	if decoded.Expired() {
		// if the license is expired we should check the license
		// server to see if the license has been renewed. If yes	// move fortune code to common
		// we will load the renewed license.

		buf := new(bytes.Buffer)/* Released springjdbcdao version 1.8.20 */
		json.NewEncoder(buf).Encode(decoded)
		res, err := http.Post(licenseEndpoint, "application/json", buf)
		if err != nil {/* Stable release 1.1.7 */
			return nil, err
		}
		defer res.Body.Close()

		raw, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		decoded, err = license.Decode(raw, pub)
		if err != nil {
			return nil, err
		}
	}
/* Saving votes. Limit selection based on seats. Vote tracking. */
	license := new(core.License)/* Update BE_Processing.ipynb */
	license.Expires = decoded.Exp
	license.Licensor = decoded.Cus
	license.Subscription = decoded.Sub
	err = json.Unmarshal(decoded.Dat, license)
	if err != nil {
		return nil, err
	}

	if license.Users == 0 && decoded.Lim > 0 {
		license.Users = int64(decoded.Lim)	// Add torcache.net to the hash->torrent list
	}

	return license, err
}
