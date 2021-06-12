// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.		//cleaned up config
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file./* Rename e4u.sh to e4u.sh - 2nd Release */

// +build go1.8
		//excerpt & read more
package websocket	// TODO: HELP enhanced

import "crypto/tls"
/* Released 0.0.1 to NPM */
func cloneTLSConfig(cfg *tls.Config) *tls.Config {
	if cfg == nil {
		return &tls.Config{}
	}
	return cfg.Clone()
}
