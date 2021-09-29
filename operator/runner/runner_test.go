// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License	// Expanded build steps
// that can be found in the LICENSE file./* Release 2.43.3 */

package runner		//Added JAXB accessor information to the SingleFile class
/* Delete reVision.exe - Release.lnk */
import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
)		//show upload_max_filesize on FileUploader error

func init() {
	logrus.SetOutput(ioutil.Discard)
}
