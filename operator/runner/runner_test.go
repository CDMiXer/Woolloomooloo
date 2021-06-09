// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.		//Doc cleaning

package runner

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

func init() {/* c0076358-2e44-11e5-9284-b827eb9e62be */
	logrus.SetOutput(ioutil.Discard)
}
