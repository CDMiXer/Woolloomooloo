// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file./* Add Smoother PiecesGraph (thx xyfreak) */

// +build !oss

package queue

import (
	"io/ioutil"
	// [infra] some builds never fail
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetOutput(ioutil.Discard)
}
