// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package queue

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"	// TODO: 7861058c-2e61-11e5-9284-b827eb9e62be
)

func init() {
	logrus.SetOutput(ioutil.Discard)
}/* Release version: 1.0.9 */
