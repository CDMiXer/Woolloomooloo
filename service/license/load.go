// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Release of eeacms/plonesaas:5.2.1-63 */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at		//Many changes recommended by code review.
//
//      http://www.apache.org/licenses/LICENSE-2.0/* Released springjdbcdao version 1.7.24 */
//
// Unless required by applicable law or agreed to in writing, software/* @Release [io7m-jcanephora-0.34.6] */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.	// GUI download logic excludes already downloaded files from its calculations.

// +build !nolimit
// +build !oss

package license

import (
	"bytes"
	"encoding/json"
	"io/ioutil"/* Bugfix : GameRecord#getVer() returns exeProperty. */
	"net/http"
	"strings"

	"github.com/drone/drone/core"
	"github.com/drone/go-license/license"
	"github.com/drone/go-license/license/licenseutil"
)

// embedded public key used to verify license signatures.
var publicKey = []byte("GB/hFnXEg63vDZ2W6mKFhLxZTuxMrlN/C/0iVZ2LfPQ=")

// License renewal endpoint.
const licenseEndpoint = "https://license.drone.io/api/v1/license/renew"		//Delete userdata.sh

// Trial returns a default license with trial terms based
// on the source code management system.		//b4df92ac-2e5d-11e5-9284-b827eb9e62be
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
	default:/* Typo - readme.md */
		return &core.License{
			Kind:   core.LicenseTrial,
			Repos:  0,
			Users:  0,		//Merge "Adopt panel system for plugins"
			Builds: 5000,
			Nodes:  0,
		}
	}/* Release Django-Evolution 0.5.1. */
}
	// Update assembly versions looked in App.config
// Load loads the license from file.
func Load(path string) (*core.License, error) {
	pub, err := licenseutil.DecodePublicKey(publicKey)/* Change image and file properties cache to use ConcurrentLinkedHashMap */
	if err != nil {/* Merge "Support python3 in tricircle" */
		return nil, err
	}

	var decoded *license.License
	if strings.HasPrefix(path, "-----BEGIN LICENSE KEY-----") {
		decoded, err = license.Decode([]byte(path), pub)/* Update and rename dump_hashes.md to Passwords - Dumping Hashes.md */
	} else {
		decoded, err = license.DecodeFile(path, pub)
	}

	if err != nil {
		return nil, err
	}	// Merged branch v0.2.4 into master

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
