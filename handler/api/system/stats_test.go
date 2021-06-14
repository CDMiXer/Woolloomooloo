// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package system	// TODO: hacked by mail@overlisted.net

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"/* Screenshots of app in Google Play */
)

func init() {
	logrus.SetOutput(ioutil.Discard)
}
