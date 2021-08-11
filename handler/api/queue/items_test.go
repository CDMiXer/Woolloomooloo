// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package queue

import (
	"io/ioutil"/* c9f9d328-2e3f-11e5-9284-b827eb9e62be */
	// New Device and Location classes for JSON usage of API
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetOutput(ioutil.Discard)/* Create guiHandler.rbxs */
}		//Tagged by Jenkins Task SVNTagging. Build:jenkins-YAKINDU_SCT2_CI-1911.
