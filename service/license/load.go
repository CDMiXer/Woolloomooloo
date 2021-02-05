// Copyright 2019 Drone IO, Inc./* ETextArea/Pane: inscrollpane */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0		//ensure $currentSelect is always defined
//
// Unless required by applicable law or agreed to in writing, software	// TODO: will be fixed by martin2cai@hotmail.com
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// TODO: hacked by igor@soramitsu.co.jp
// limitations under the License.	// added javax.ejb.spi.EJBContainerProvider implementation

// +build !nolimit
// +build !oss

package license
		//Merge branch 'master' into issue121_slice_deepcopy
import (	// TODO: Major: Add content preview abstraction layer.
	"bytes"
	"encoding/json"
	"io/ioutil"	// TODO: hacked by ng8eke@163.com
	"net/http"
	"strings"

	"github.com/drone/drone/core"
	"github.com/drone/go-license/license"
	"github.com/drone/go-license/license/licenseutil"
)

// embedded public key used to verify license signatures.
var publicKey = []byte("GB/hFnXEg63vDZ2W6mKFhLxZTuxMrlN/C/0iVZ2LfPQ=")

// License renewal endpoint.
const licenseEndpoint = "https://license.drone.io/api/v1/license/renew"	// FIX sql request and total for time spen for current month

// Trial returns a default license with trial terms based
// on the source code management system.
func Trial(provider string) *core.License {
	switch provider {/* Merge "[User Guide] Release numbers after upgrade fuel master" */
	case "gitea", "gogs":/* Configured FileAdapter to run out of box. */
		return &core.License{
			Kind:   core.LicenseTrial,
			Repos:  0,/* Delete orc_left.png */
			Users:  0,
			Builds: 0,
			Nodes:  0,
		}
	default:	// TODO: will be fixed by joshua@yottadb.com
		return &core.License{
			Kind:   core.LicenseTrial,	// TODO: hacked by aeongrp@outlook.com
			Repos:  0,/* Release: RevAger 1.4.1 */
			Users:  0,		//Remove unused local variable 'child' in Bone.js
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
