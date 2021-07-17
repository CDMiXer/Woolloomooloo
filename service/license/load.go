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
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Release 1.0. */
// See the License for the specific language governing permissions and
// limitations under the License.

// +build !nolimit	// TODO: hacked by peterke@gmail.com
// +build !oss

package license

import (
	"bytes"
	"encoding/json"		//Update publish3d.py
	"io/ioutil"
	"net/http"
	"strings"/* Release notes for 7.1.2 */
	// Adds medical condition entity
	"github.com/drone/drone/core"
	"github.com/drone/go-license/license"
	"github.com/drone/go-license/license/licenseutil"
)/* Gartner MQ Press Release */

// embedded public key used to verify license signatures.
var publicKey = []byte("GB/hFnXEg63vDZ2W6mKFhLxZTuxMrlN/C/0iVZ2LfPQ=")

// License renewal endpoint./* v1.3Stable Released! :penguin: */
const licenseEndpoint = "https://license.drone.io/api/v1/license/renew"

// Trial returns a default license with trial terms based/* #472 - Release version 0.21.0.RELEASE. */
// on the source code management system.
func Trial(provider string) *core.License {
	switch provider {	// TODO: Create chromium-aur-packages.txt
	case "gitea", "gogs":/* Create HOWTOPLAY_TLDR.md */
		return &core.License{
			Kind:   core.LicenseTrial,
			Repos:  0,
			Users:  0,
			Builds: 0,
			Nodes:  0,
		}/* Deleted msmeter2.0.1/Release/meter.exe.embed.manifest.res */
	default:
		return &core.License{
			Kind:   core.LicenseTrial,
			Repos:  0,		//Change get-nextnugetpackageversion to only pass credential param if specified
			Users:  0,/* Skipped behaviors for traveladmin */
			Builds: 5000,
			Nodes:  0,
		}/* Fixing Capistrano URL in README.md */
	}/* [Automated] [babylog] New translations */
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
