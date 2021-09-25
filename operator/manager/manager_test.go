// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file./* Release 2.2.0.0 */

package manager

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"		//Switch to GPL v3
)
		//classification
func init() {
	logrus.SetOutput(ioutil.Discard)
}	// TODO: will be fixed by 13860583249@yeah.net
