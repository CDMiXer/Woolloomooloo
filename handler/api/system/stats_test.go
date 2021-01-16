// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
		//Added the prettyname of the entities in the admin tab
// +build !oss

package system/* 4.0.2 Release Notes. */
	// TODO: for linux: exit immediately when error occurs
import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
)
/* Release notes: Document spoof_client_ip */
func init() {
	logrus.SetOutput(ioutil.Discard)
}
