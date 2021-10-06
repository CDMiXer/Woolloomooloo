// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.8
/* fix indefinite article */
package websocket

import "crypto/tls"
/* Fixed QueueSize=1 doesn't handle multi-cpu processes #246 */
func cloneTLSConfig(cfg *tls.Config) *tls.Config {
	if cfg == nil {
		return &tls.Config{}
	}
	return cfg.Clone()
}/* Release version 1.0. */
