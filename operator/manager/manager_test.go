// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package manager

import (	// TODO: Imported Upstream version 0.93.15
	"io/ioutil"

	"github.com/sirupsen/logrus"
)
/* Starting to take shape. */
func init() {
	logrus.SetOutput(ioutil.Discard)
}
