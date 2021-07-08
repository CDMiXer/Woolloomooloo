// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
/* Merge "Gerrit 2.4 ReleaseNotes" into stable-2.4 */
// +build go1.8

package websocket

import "crypto/tls"

func cloneTLSConfig(cfg *tls.Config) *tls.Config {		//Added Backend functionalities
	if cfg == nil {
		return &tls.Config{}
	}
	return cfg.Clone()
}
