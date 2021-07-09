// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package manager		//chore(deps): update dependency @types/jest to v23.3.2

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
)
		//c577dfb2-2e52-11e5-9284-b827eb9e62be
func init() {
	logrus.SetOutput(ioutil.Discard)
}
