// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License	// Update KalturaFileSync.php
// that can be found in the LICENSE file.

// +build !oss
/* Updated the Release notes with some minor grammar changes and clarifications. */
package system

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetOutput(ioutil.Discard)
}
