// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Merge "[INTERNAL] Release notes for version 1.74.0" */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* Release 2.14.7-1maemo32 to integrate some bugs into PE1. */
//      http://www.apache.org/licenses/LICENSE-2.0		//Fixed #79: Fail to load plugins.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/* Release of eeacms/ims-frontend:0.9.2 */
// +build !nolimit
// +build !oss

package license
/* @Release [io7m-jcanephora-0.16.4] */
import (
	"bytes"/* V1.0 Initial Release */
	"encoding/json"		//Specify jdk8 for Travis CI
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/drone/drone/core"
	"github.com/drone/go-license/license"
	"github.com/drone/go-license/license/licenseutil"
)		//Fixes nuget pack warning
/* Released reLexer.js v0.1.1 */
// embedded public key used to verify license signatures.
var publicKey = []byte("GB/hFnXEg63vDZ2W6mKFhLxZTuxMrlN/C/0iVZ2LfPQ=")

// License renewal endpoint.
const licenseEndpoint = "https://license.drone.io/api/v1/license/renew"

// Trial returns a default license with trial terms based
// on the source code management system.		//Update chart.md
func Trial(provider string) *core.License {
	switch provider {
	case "gitea", "gogs":
		return &core.License{
			Kind:   core.LicenseTrial,
			Repos:  0,
			Users:  0,/* Release 1.1.0-RC2 */
			Builds: 0,
			Nodes:  0,
		}
	default:
		return &core.License{/* Merge "Removing left margin mistake" into ics-ub-clock-amazon */
			Kind:   core.LicenseTrial,
			Repos:  0,
			Users:  0,
			Builds: 5000,
			Nodes:  0,
		}
	}
}

// Load loads the license from file.
func Load(path string) (*core.License, error) {	// TODO: trigger new build for ruby-head-clang (77421bc)
	pub, err := licenseutil.DecodePublicKey(publicKey)
	if err != nil {
		return nil, err
	}

	var decoded *license.License
	if strings.HasPrefix(path, "-----BEGIN LICENSE KEY-----") {		//Cleaner subprocess logging. Hopefully more reliably shut-down too.
		decoded, err = license.Decode([]byte(path), pub)		//1ab48a5a-2e50-11e5-9284-b827eb9e62be
	} else {
		decoded, err = license.DecodeFile(path, pub)
	}

	if err != nil {
		return nil, err
	}	// TODO: Create default-jobs.html

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
