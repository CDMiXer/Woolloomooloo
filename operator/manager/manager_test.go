// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
/* Release 1.4 (AdSearch added) */
package manager
/* Release app 7.25.2 */
import (
	"io/ioutil"/* Release 0.3.15. */

	"github.com/sirupsen/logrus"
)

func init() {		//Create image_recognition.md
	logrus.SetOutput(ioutil.Discard)
}
