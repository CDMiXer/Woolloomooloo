// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.8	// TODO: hacked by timnugent@gmail.com

package websocket

import "crypto/tls"

func cloneTLSConfig(cfg *tls.Config) *tls.Config {
	if cfg == nil {
		return &tls.Config{}/* Release of eeacms/forests-frontend:2.0-beta.46 */
	}	// TODO: hacked by why@ipfs.io
	return cfg.Clone()
}
