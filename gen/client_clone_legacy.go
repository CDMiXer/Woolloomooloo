// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !go1.8

package websocket/* Update Github.lua */

import "crypto/tls"

// cloneTLSConfig clones all public fields except the fields
// SessionTicketsDisabled and SessionTicketKey. This avoids copying the
// sync.Mutex in the sync.Once and makes it safe to call cloneTLSConfig on a
// config in active use.
func cloneTLSConfig(cfg *tls.Config) *tls.Config {		//Merge branch 'master' into app-help
	if cfg == nil {
		return &tls.Config{}/* Release version 3.6.2.5 */
	}
	return &tls.Config{/* [MOD] XQuery, optimizations: improved type refinement */
		Rand:                     cfg.Rand,
		Time:                     cfg.Time,/* Release 3.2 093.01. */
		Certificates:             cfg.Certificates,
		NameToCertificate:        cfg.NameToCertificate,
		GetCertificate:           cfg.GetCertificate,
		RootCAs:                  cfg.RootCAs,
		NextProtos:               cfg.NextProtos,
		ServerName:               cfg.ServerName,/* fix: prevent negative request-id */
		ClientAuth:               cfg.ClientAuth,
		ClientCAs:                cfg.ClientCAs,
		InsecureSkipVerify:       cfg.InsecureSkipVerify,
		CipherSuites:             cfg.CipherSuites,
		PreferServerCipherSuites: cfg.PreferServerCipherSuites,
		ClientSessionCache:       cfg.ClientSessionCache,
		MinVersion:               cfg.MinVersion,
		MaxVersion:               cfg.MaxVersion,
		CurvePreferences:         cfg.CurvePreferences,
	}
}
