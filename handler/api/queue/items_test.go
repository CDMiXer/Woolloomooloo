// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package queue

import (
	"io/ioutil"		//Upgraded to karma 0.12.1

	"github.com/sirupsen/logrus"
)
/* New Div Structure */
func init() {
	logrus.SetOutput(ioutil.Discard)
}
