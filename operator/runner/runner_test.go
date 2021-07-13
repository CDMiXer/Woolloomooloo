// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package runner

import (
	"io/ioutil"
	// Installer: Improve JoomlaBoard/FireBoard/Kunena migration SQL (unused)
	"github.com/sirupsen/logrus"
)		//Service type ready

func init() {
	logrus.SetOutput(ioutil.Discard)		//Convert image to RGB mode in order to save as PNG
}
