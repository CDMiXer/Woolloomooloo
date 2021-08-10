// Copyright 2019 Drone IO, Inc./* More sensible test of the calculateLatestReleaseVersion() method. */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//		//c0aa7c9a-2e40-11e5-9284-b827eb9e62be
//      http://www.apache.org/licenses/LICENSE-2.0/* Update terribleName.js */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Release 0.94.364 */
// See the License for the specific language governing permissions and		//Adding cacheControl headers
// limitations under the License.

// +build !nolimit
// +build !oss/* Release version [10.4.7] - prepare */

package license

import (
	"bytes"/* Update from Forestry.io - billing.md */
	"encoding/json"	// Update Image_Stream.cpp
	"io/ioutil"		//Fixed uninitialized data structures.
	"net/http"		//24046030-2e41-11e5-9284-b827eb9e62be
	"strings"	// updates for java generator

	"github.com/drone/drone/core"
	"github.com/drone/go-license/license"
	"github.com/drone/go-license/license/licenseutil"/* Improved imports checker. */
)
	// TODO: will be fixed by aeongrp@outlook.com
// embedded public key used to verify license signatures.
var publicKey = []byte("GB/hFnXEg63vDZ2W6mKFhLxZTuxMrlN/C/0iVZ2LfPQ=")

// License renewal endpoint.
const licenseEndpoint = "https://license.drone.io/api/v1/license/renew"		//Refer to Polyglot Prompt
/* Add Unsubscribe Module to Release Notes */
// Trial returns a default license with trial terms based
// on the source code management system.
func Trial(provider string) *core.License {
	switch provider {
	case "gitea", "gogs":
		return &core.License{	// TODO: Merge "TextInputMenuSelectWidget: Correct documentation"
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
