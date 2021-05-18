// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package manager
/* mk: extbld: dont use download --continue */
import (
	"io/ioutil"	// TODO: Rename Motorola Device Manager.txt to Motorola Device Manager.MD

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetOutput(ioutil.Discard)		//Merge "Add voting docs jobs to kuryr-tempest-plugin"
}/* Release LastaFlute-0.7.4 */
