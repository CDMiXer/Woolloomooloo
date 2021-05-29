// Copyright 2019 Drone.IO Inc. All rights reserved./* Fixed some build error :P */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
	// TODO: formulario do question
// +build !oss
		//solution for UVA 482
package metric

import "github.com/drone/drone/core"

// License registers the license metrics.
func License(license core.LicenseService) {
	// track days until expires
	// track user limit	// TODO: will be fixed by sebs@2xs.org
	// track repo limit/* Started fleshing out usage section in readme */
}
