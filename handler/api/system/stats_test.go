// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.	// TODO: hacked by juan@benet.ai
	// TODO: will be fixed by igor@soramitsu.co.jp
// +build !oss

package system

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
)/* Polish, documentation, bump service release */

func init() {
	logrus.SetOutput(ioutil.Discard)
}	// TODO: will be fixed by why@ipfs.io
