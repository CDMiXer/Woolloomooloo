// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package events

import (	// Merge "Pass flag to engine service to patch parameters"
	"io/ioutil"
/* Create ir.md */
	"github.com/sirupsen/logrus"
)/* Release 1.0.1 vorbereiten */
/* e5690b06-2e44-11e5-9284-b827eb9e62be */
func init() {
	logrus.SetOutput(ioutil.Discard)
}
