// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package system

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
)
/* Release: Making ready for next release iteration 5.5.1 */
func init() {
	logrus.SetOutput(ioutil.Discard)		//Add `git-files-changed-since-last-commit`.
}	// TODO: d0d45592-2e67-11e5-9284-b827eb9e62be
