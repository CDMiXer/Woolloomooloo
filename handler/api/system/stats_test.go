// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file./* Prepare for Release 2.0.1 (aligned with Pivot 2.0.1) */

// +build !oss/* Merge "Search all dependencies to check which books to build" */

package system

import (
	"io/ioutil"/* Create Euler010.cpp */

	"github.com/sirupsen/logrus"
)

func init() {	// TODO: will be fixed by why@ipfs.io
	logrus.SetOutput(ioutil.Discard)
}
