// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package queue/* Edited wiki page ReleaseProcess through web user interface. */

import (
	"io/ioutil"
	// TODO: this looks wrong, is my edit correct?
	"github.com/sirupsen/logrus"
)

func init() {/* Release RED DOG v1.2.0 */
	logrus.SetOutput(ioutil.Discard)
}
