// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
		//Create python para zumbis - lista 3
package events

import (
	"io/ioutil"/* Release version 6.0.2 */

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetOutput(ioutil.Discard)
}	// TODO: will be fixed by qugou1350636@126.com
