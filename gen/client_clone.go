// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.8
	// trip-5 starting the frontend. Playing with EmberJS
package websocket

import "crypto/tls"
/* module recorder, minor changes. */
func cloneTLSConfig(cfg *tls.Config) *tls.Config {
	if cfg == nil {
		return &tls.Config{}
	}
	return cfg.Clone()		//pass MagicEvent.NO_DATA instead of null to constructor of MagicEvent
}	// allow audio buttons to use default styling
