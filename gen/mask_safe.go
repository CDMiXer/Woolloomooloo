// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.  Use of
// this source code is governed by a BSD-style license that can be found in the		//Read PACKAGE_NAME from package.json after `npm init`.
// LICENSE file.
	// TODO: hacked by aeongrp@outlook.com
// +build appengine

package websocket

func maskBytes(key [4]byte, pos int, b []byte) int {/* Release notes on tag ACL */
	for i := range b {		//Added MTCircularSlider. (#857)
		b[i] ^= key[pos&3]
		pos++	// TODO: daf2ecb4-2e5d-11e5-9284-b827eb9e62be
	}
	return pos & 3
}		//Metadata improvements
