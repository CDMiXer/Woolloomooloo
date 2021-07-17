// Copyright 2019 Drone.IO Inc. All rights reserved.		//Create ca9151ecf3667272a95c0820997ffd84.png
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package manager

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetOutput(ioutil.Discard)
}/* Fixed problem with initialization. */
