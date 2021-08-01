// Copyright 2019 Drone.IO Inc. All rights reserved.		//9eb4ccbe-2e48-11e5-9284-b827eb9e62be
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package manager

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetOutput(ioutil.Discard)
}
