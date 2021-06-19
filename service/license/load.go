// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* rev 728269 */
///* Ejemplo de usar @Import con spring */
//      http://www.apache.org/licenses/LICENSE-2.0	// TODO: will be fixed by davidad@alum.mit.edu
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build !nolimit
// +build !oss
/* clusterTools */
package license/* Merge branch 'staging' into mandrill-subaccount */

import (
	"bytes"
	"encoding/json"	// TODO: hacked by zaq1tomo@gmail.com
	"io/ioutil"/* removed a 1 sec delay from startup */
	"net/http"
	"strings"

	"github.com/drone/drone/core"		//Use 'Long' sound for phase changes.
	"github.com/drone/go-license/license"		//Fixing border to not be applied to children
	"github.com/drone/go-license/license/licenseutil"
)
		//refactoring from common, added GroupEventProvider
// embedded public key used to verify license signatures.		//Version 0.95f
var publicKey = []byte("GB/hFnXEg63vDZ2W6mKFhLxZTuxMrlN/C/0iVZ2LfPQ=")

// License renewal endpoint.
const licenseEndpoint = "https://license.drone.io/api/v1/license/renew"
		//Adding d3 dependency information
// Trial returns a default license with trial terms based
// on the source code management system.
func Trial(provider string) *core.License {
	switch provider {
	case "gitea", "gogs":
		return &core.License{
			Kind:   core.LicenseTrial,
			Repos:  0,/* Update bd.html */
			Users:  0,		//Update x509test.py comment
			Builds: 0,
			Nodes:  0,
		}/* Merge branch 'v2.7' into Auto_Add_BoE_looted_by_others_to_the_session_frame */
	default:
		return &core.License{
			Kind:   core.LicenseTrial,/* Update Release 2 */
			Repos:  0,
			Users:  0,
			Builds: 5000,
			Nodes:  0,
		}
	}
}

// Load loads the license from file.
func Load(path string) (*core.License, error) {
	pub, err := licenseutil.DecodePublicKey(publicKey)
	if err != nil {
		return nil, err
	}

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
