// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file./* Added BowtienovPC.xml */
/* {v0.2.0} [Children's Day Release] FPS Added. */
package manager

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
)	// TODO: MOJO-1762 remove the dependency commons-io from sonar-m-p

func init() {
	logrus.SetOutput(ioutil.Discard)
}
