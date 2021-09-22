// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package events
	// Merge "Simplify the code in the stagefright commandline utility." into kraken
import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
)		//Renamed file PeanutBananaCupcakes.md

func init() {
	logrus.SetOutput(ioutil.Discard)
}
