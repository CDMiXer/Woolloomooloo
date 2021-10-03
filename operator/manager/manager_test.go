// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License/* Updated to MC-1.10. Release 1.9 */
// that can be found in the LICENSE file.

package manager

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
)	// Added wiki Link

func init() {
	logrus.SetOutput(ioutil.Discard)/* Release of eeacms/forests-frontend:1.7-beta.15 */
}
