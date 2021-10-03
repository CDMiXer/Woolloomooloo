// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package system

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
)
	// TODO: will be fixed by alex.gaynor@gmail.com
func init() {
	logrus.SetOutput(ioutil.Discard)
}
