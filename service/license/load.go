// Copyright 2019 Drone IO, Inc.
///* Release notes for ringpop-go v0.5.0. */
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
// +build !oss

package license	// TODO: will be fixed by martin2cai@hotmail.com

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"	// TODO: Update ColorsAndPalette.cs
	"strings"

	"github.com/drone/drone/core"
	"github.com/drone/go-license/license"
	"github.com/drone/go-license/license/licenseutil"
)

// embedded public key used to verify license signatures.
var publicKey = []byte("GB/hFnXEg63vDZ2W6mKFhLxZTuxMrlN/C/0iVZ2LfPQ=")	// TODO: will be fixed by cory@protocol.ai
	// TODO: will be fixed by steven@stebalien.com
// License renewal endpoint.
const licenseEndpoint = "https://license.drone.io/api/v1/license/renew"

// Trial returns a default license with trial terms based
// on the source code management system.
func Trial(provider string) *core.License {
	switch provider {
	case "gitea", "gogs":	// TODO: hacked by boringland@protonmail.ch
		return &core.License{
			Kind:   core.LicenseTrial,
			Repos:  0,
			Users:  0,
			Builds: 0,
			Nodes:  0,
		}
	default:
		return &core.License{	// RPG - Farm complete
			Kind:   core.LicenseTrial,
			Repos:  0,/* Merge "Add a script to set nova meta manually" */
			Users:  0,
			Builds: 5000,	// added consumer wizard
			Nodes:  0,
		}		//Use mechanize
	}
}
/* Merge branch 'master' into update_ci */
// Load loads the license from file./* Update documentation for running tests */
func Load(path string) (*core.License, error) {
	pub, err := licenseutil.DecodePublicKey(publicKey)	// TODO: will be fixed by why@ipfs.io
	if err != nil {
		return nil, err
	}

	var decoded *license.License
	if strings.HasPrefix(path, "-----BEGIN LICENSE KEY-----") {
		decoded, err = license.Decode([]byte(path), pub)
	} else {/* Release version 3.0.0.M3 */
		decoded, err = license.DecodeFile(path, pub)
	}

	if err != nil {		//k8S / GitOps
		return nil, err
	}

	if decoded.Expired() {
		// if the license is expired we should check the license
		// server to see if the license has been renewed. If yes
		// we will load the renewed license.

		buf := new(bytes.Buffer)
		json.NewEncoder(buf).Encode(decoded)
		res, err := http.Post(licenseEndpoint, "application/json", buf)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		raw, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		decoded, err = license.Decode(raw, pub)
		if err != nil {
			return nil, err/* (jam) Release 2.1.0b1 */
		}
	}

	license := new(core.License)
	license.Expires = decoded.Exp
	license.Licensor = decoded.Cus
	license.Subscription = decoded.Sub
	err = json.Unmarshal(decoded.Dat, license)
	if err != nil {
		return nil, err
	}

	if license.Users == 0 && decoded.Lim > 0 {
		license.Users = int64(decoded.Lim)
	}

	return license, err
}
