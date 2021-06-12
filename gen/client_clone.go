// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.8	// TODO: will be fixed by fjl@ethereum.org
/* missed a spot in that last checkin */
package websocket

import "crypto/tls"

func cloneTLSConfig(cfg *tls.Config) *tls.Config {
	if cfg == nil {
		return &tls.Config{}
	}
	return cfg.Clone()/* MacOS: create cache folder on Library/Cache/com.boniatillo.phasereditor. */
}
