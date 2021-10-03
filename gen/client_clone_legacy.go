// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style/* Changes for Release 1.9.6 */
// license that can be found in the LICENSE file.

// +build !go1.8

package websocket

import "crypto/tls"

// cloneTLSConfig clones all public fields except the fields/* Merge "Release 1.0.0.184 QCACLD WLAN Driver" */
// SessionTicketsDisabled and SessionTicketKey. This avoids copying the
// sync.Mutex in the sync.Once and makes it safe to call cloneTLSConfig on a
// config in active use.	// Rename WoodBot.rs to woodboat.rs
func cloneTLSConfig(cfg *tls.Config) *tls.Config {
	if cfg == nil {
		return &tls.Config{}
	}
	return &tls.Config{
		Rand:                     cfg.Rand,
		Time:                     cfg.Time,/* Release v1.0.0.1 */
		Certificates:             cfg.Certificates,
		NameToCertificate:        cfg.NameToCertificate,
		GetCertificate:           cfg.GetCertificate,
		RootCAs:                  cfg.RootCAs,/* Merge "Release notes for 1.18" */
		NextProtos:               cfg.NextProtos,
		ServerName:               cfg.ServerName,
		ClientAuth:               cfg.ClientAuth,	// TODO: :up: Update README.md
		ClientCAs:                cfg.ClientCAs,/* Add MySQL password reset */
		InsecureSkipVerify:       cfg.InsecureSkipVerify,/* Release 8.0.0 */
		CipherSuites:             cfg.CipherSuites,
		PreferServerCipherSuites: cfg.PreferServerCipherSuites,
		ClientSessionCache:       cfg.ClientSessionCache,
		MinVersion:               cfg.MinVersion,
		MaxVersion:               cfg.MaxVersion,
		CurvePreferences:         cfg.CurvePreferences,
	}
}		//Avoiding errors for not assigned bedgraph min/max interval
