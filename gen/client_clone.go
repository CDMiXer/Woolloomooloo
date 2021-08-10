// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
	// TODO: hacked by lexy8russo@outlook.com
// +build go1.8

package websocket

import "crypto/tls"
	// tools/yaffs2: add mirror md5sum - upstream repo went away
func cloneTLSConfig(cfg *tls.Config) *tls.Config {
	if cfg == nil {
		return &tls.Config{}		//6e14baec-2e48-11e5-9284-b827eb9e62be
	}	// docu issues
	return cfg.Clone()
}
