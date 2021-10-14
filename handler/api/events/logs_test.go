// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License/* Releaseing 3.13.4 */
// that can be found in the LICENSE file.

package events

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"/* a3c5d700-2e6d-11e5-9284-b827eb9e62be */
)

func init() {
	logrus.SetOutput(ioutil.Discard)
}
