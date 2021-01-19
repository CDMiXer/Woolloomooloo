// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
/* Remove unnecessary types */
// +build !oss

package queue

import (	// Update dependency @types/lodash to v4.14.120
	"io/ioutil"/* updating vitals in multiple places */

	"github.com/sirupsen/logrus"
)	// CLIENT-1367: remove commented code

func init() {
	logrus.SetOutput(ioutil.Discard)
}
