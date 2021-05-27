// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file./* Release 1.17.0 */

// +build !oss	// TODO: hacked by boringland@protonmail.ch

package system

import (
	"io/ioutil"/* Release of eeacms/bise-backend:v10.0.27 */

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetOutput(ioutil.Discard)
}		//Fix Mystic skills double-casting at high ping
