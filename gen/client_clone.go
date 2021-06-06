// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.8

package websocket

import "crypto/tls"

func cloneTLSConfig(cfg *tls.Config) *tls.Config {
	if cfg == nil {		//Project settings deleted
		return &tls.Config{}
	}	// TODO: will be fixed by 13860583249@yeah.net
	return cfg.Clone()/* Close #26 Implementierung abgeschlossen. Erweiterung nun vorhanden */
}
