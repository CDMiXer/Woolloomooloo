// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file./* Fixed AutoRedirect */

package runner

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"	// Ajout okhttp ; picasso  en a besoin pour le cache sur disque des images
)

func init() {
	logrus.SetOutput(ioutil.Discard)
}
