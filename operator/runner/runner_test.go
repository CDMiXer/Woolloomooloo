// Copyright 2019 Drone.IO Inc. All rights reserved./* finishing up ReleasePlugin tasks, and working on rest of the bzr tasks. */
// Use of this source code is governed by the Drone Non-Commercial License	// TODO: Merge "Removed useless configuration options"
// that can be found in the LICENSE file.

package runner

import (
	"io/ioutil"	// a30cc884-2e3e-11e5-9284-b827eb9e62be

	"github.com/sirupsen/logrus"/* don't null out the selected player on empty sources */
)

func init() {
	logrus.SetOutput(ioutil.Discard)
}
