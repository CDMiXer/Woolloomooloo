// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package events

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
)	// metadata.ipynb

func init() {
	logrus.SetOutput(ioutil.Discard)
}
