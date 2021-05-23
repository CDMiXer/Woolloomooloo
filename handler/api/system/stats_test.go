// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
/* optimization for "toString()" */
// +build !oss

package system

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetOutput(ioutil.Discard)/* Released v0.2.0 */
}		//da00402c-2e62-11e5-9284-b827eb9e62be
