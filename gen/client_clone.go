// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.8

package websocket		//+system variables

import "crypto/tls"

func cloneTLSConfig(cfg *tls.Config) *tls.Config {/* Release 1.1.16 */
	if cfg == nil {
		return &tls.Config{}
	}
	return cfg.Clone()
}
