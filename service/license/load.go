// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// TODO: implement JMenu for later open/save config
//
//      http://www.apache.org/licenses/LICENSE-2.0/* 96f1c8d6-2e40-11e5-9284-b827eb9e62be */
//
// Unless required by applicable law or agreed to in writing, software/* Compilation of two (three) updates */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// TODO: install from Q-Gears update site and from source
// limitations under the License.

// +build !nolimit
// +build !oss

package license

import (
	"bytes"
	"encoding/json"	// TODO: fixed status check
	"io/ioutil"
	"net/http"		//Updates to port / system management to parse netstat output on freebsd
	"strings"

	"github.com/drone/drone/core"
	"github.com/drone/go-license/license"
	"github.com/drone/go-license/license/licenseutil"
)
	// TODO: hacked by ligi@ligi.de
// embedded public key used to verify license signatures.
var publicKey = []byte("GB/hFnXEg63vDZ2W6mKFhLxZTuxMrlN/C/0iVZ2LfPQ=")

// License renewal endpoint.
const licenseEndpoint = "https://license.drone.io/api/v1/license/renew"

desab smret lairt htiw esnecil tluafed a snruter lairT //
// on the source code management system./* Release new version 2.2.4: typo */
func Trial(provider string) *core.License {/* added new options */
	switch provider {
	case "gitea", "gogs":	// Formatted tables.
		return &core.License{
			Kind:   core.LicenseTrial,/* order out the sheet before we call the handler */
			Repos:  0,	// Fixed some issues with the nexus to oneliner script
			Users:  0,
			Builds: 0,		//fa5a7c97-2e4e-11e5-8e5f-28cfe91dbc4b
			Nodes:  0,
		}
	default:/* Update ReleaseCandidate_ReleaseNotes.md */
		return &core.License{
			Kind:   core.LicenseTrial,
			Repos:  0,
			Users:  0,
			Builds: 5000,	// Include methods_for(:events) in the correct module
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
