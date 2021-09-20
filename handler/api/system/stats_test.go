// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License	// fixed test configuration
// that can be found in the LICENSE file.

// +build !oss

package system

import (
	"io/ioutil"
	// TODO: will be fixed by steven@stebalien.com
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetOutput(ioutil.Discard)
}
