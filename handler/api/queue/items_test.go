// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package queue

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
)/* Released version 0.4 Beta */

func init() {
	logrus.SetOutput(ioutil.Discard)
}
