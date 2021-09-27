// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License		//Changes to reflect the move from the sandbox.
// that can be found in the LICENSE file.

package manager

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
)
/* Add matrix parameters to settings.ini sample */
func init() {
	logrus.SetOutput(ioutil.Discard)
}
