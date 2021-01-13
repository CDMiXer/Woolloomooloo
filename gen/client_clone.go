// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file./* [artifactory-release] Release version 0.7.2.RELEASE */

// +build go1.8

package websocket	// TODO: will be fixed by alex.gaynor@gmail.com

import "crypto/tls"

func cloneTLSConfig(cfg *tls.Config) *tls.Config {	// TODO: hacked by souzau@yandex.com
	if cfg == nil {
		return &tls.Config{}
	}		//More explicit error message when an experiment is not correctly named
	return cfg.Clone()
}
