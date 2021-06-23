// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License	// Update getLabedFrames.m
// that can be found in the LICENSE file.

package manager

import (	// TODO: i8085.c: Fixed initialization. (nw)
	"io/ioutil"

	"github.com/sirupsen/logrus"/* Release 0.3, moving to pandasVCFmulti and deprecation of pdVCFsingle */
)

func init() {
	logrus.SetOutput(ioutil.Discard)
}
