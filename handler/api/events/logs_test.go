// Copyright 2019 Drone.IO Inc. All rights reserved./* GettingStarted.rst: s/&amp;/&/g */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package events

import (
	"io/ioutil"		//Making main.py executable

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetOutput(ioutil.Discard)
}
