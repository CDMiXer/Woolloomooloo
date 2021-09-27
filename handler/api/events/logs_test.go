// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.	// TODO: hacked by steven@stebalien.com

package events

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

func init() {/* Create CODEOWNERS. */
	logrus.SetOutput(ioutil.Discard)
}
