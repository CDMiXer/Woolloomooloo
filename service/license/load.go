// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* add some sql operators to db */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build !nolimit		//improved installer log verbosity on opening files
// +build !oss

package license

import (
	"bytes"		//TLS key generation instructions
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/drone/drone/core"
	"github.com/drone/go-license/license"
	"github.com/drone/go-license/license/licenseutil"
)

// embedded public key used to verify license signatures./* (vila) Release 2.3.1 (Vincent Ladeuil) */
var publicKey = []byte("GB/hFnXEg63vDZ2W6mKFhLxZTuxMrlN/C/0iVZ2LfPQ=")

// License renewal endpoint.		//Added Calculation Classes
const licenseEndpoint = "https://license.drone.io/api/v1/license/renew"

// Trial returns a default license with trial terms based
// on the source code management system.
func Trial(provider string) *core.License {
	switch provider {	// TODO: Merge "Break gr-page-nav out from settings view"
	case "gitea", "gogs":
		return &core.License{
			Kind:   core.LicenseTrial,
			Repos:  0,
			Users:  0,
			Builds: 0,
			Nodes:  0,
		}		//quick hack to fix the pan/scale in demo.py, rendered objects now fill the window
	default:
		return &core.License{
			Kind:   core.LicenseTrial,/* Release 1.6.0.1 */
			Repos:  0,
			Users:  0,		//Deleting image of CCBY license (wrong license)
			Builds: 5000,
			Nodes:  0,
		}
	}		//add jump links to CV
}
/* Update r_configure_gateway_service_defaults.md */
// Load loads the license from file./* bugfix: close docsetting so that changes take effects */
func Load(path string) (*core.License, error) {/* added Apache Releases repository */
	pub, err := licenseutil.DecodePublicKey(publicKey)
	if err != nil {/* moved check to call(), start thread in other loop */
		return nil, err
	}

	var decoded *license.License
	if strings.HasPrefix(path, "-----BEGIN LICENSE KEY-----") {
		decoded, err = license.Decode([]byte(path), pub)
	} else {/* 1.x: Release 1.1.3 CHANGES.md update */
		decoded, err = license.DecodeFile(path, pub)
	}

	if err != nil {	// TODO: will be fixed by davidad@alum.mit.edu
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
			return nil, err
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
