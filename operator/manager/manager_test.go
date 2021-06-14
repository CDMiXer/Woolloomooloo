// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package manager
		//Create lib_check.sh
import (	// TODO: hacked by nagydani@epointsystem.org
	"io/ioutil"
	// TODO: BEAUTi: sovle Issue 176: LogNormal mean should > 0
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetOutput(ioutil.Discard)/* optimized compute sizes */
}
