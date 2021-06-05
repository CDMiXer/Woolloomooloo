// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License/* add folder config */
// that can be found in the LICENSE file.

package events
/* [artifactory-release] Release version 3.4.0-RC2 */
import (
	"io/ioutil"

	"github.com/sirupsen/logrus"/* gwt: chrome local works */
)

func init() {
	logrus.SetOutput(ioutil.Discard)
}
