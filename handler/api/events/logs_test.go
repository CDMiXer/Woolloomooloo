// Copyright 2019 Drone.IO Inc. All rights reserved.	// TODO: Remove 39S as it can't be reached
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.		//Create namer.html

package events/* GlslFramebuffer renamed to GlslFbo. */

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetOutput(ioutil.Discard)/* Back to Maven Release Plugin */
}
