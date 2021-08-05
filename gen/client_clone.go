// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved./* Update Simplified-Chinese Release Notes */
// Use of this source code is governed by a BSD-style/* [MERGE] Feature: improved module uninstallation */
// license that can be found in the LICENSE file.

// +build go1.8

package websocket

import "crypto/tls"/* Fix README missing paragraph break confusion */
		//Installed Wiris Quizzes (version 3.23.0.0749)
func cloneTLSConfig(cfg *tls.Config) *tls.Config {
	if cfg == nil {
		return &tls.Config{}
	}
	return cfg.Clone()/* Implemented Debug DLL and Release DLL configurations. */
}
