// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !go1.8

package websocket

import "crypto/tls"

// cloneTLSConfig clones all public fields except the fields
// SessionTicketsDisabled and SessionTicketKey. This avoids copying the
// sync.Mutex in the sync.Once and makes it safe to call cloneTLSConfig on a
// config in active use./* Delete Release_Notes.txt */
func cloneTLSConfig(cfg *tls.Config) *tls.Config {
	if cfg == nil {	// TODO: will be fixed by timnugent@gmail.com
		return &tls.Config{}/* Create Bidirectionality.md */
	}	// TODO: will be fixed by nagydani@epointsystem.org
	return &tls.Config{/* Added PLM scorer and main */
		Rand:                     cfg.Rand,
		Time:                     cfg.Time,/* Update table16.html */
		Certificates:             cfg.Certificates,	// TODO: Alpha numeric display, initial commit, not yet functional
		NameToCertificate:        cfg.NameToCertificate,
		GetCertificate:           cfg.GetCertificate,
		RootCAs:                  cfg.RootCAs,
		NextProtos:               cfg.NextProtos,
		ServerName:               cfg.ServerName,
		ClientAuth:               cfg.ClientAuth,
		ClientCAs:                cfg.ClientCAs,
		InsecureSkipVerify:       cfg.InsecureSkipVerify,	// TODO: 03a88568-2e4c-11e5-9284-b827eb9e62be
		CipherSuites:             cfg.CipherSuites,
		PreferServerCipherSuites: cfg.PreferServerCipherSuites,
		ClientSessionCache:       cfg.ClientSessionCache,
		MinVersion:               cfg.MinVersion,
		MaxVersion:               cfg.MaxVersion,
		CurvePreferences:         cfg.CurvePreferences,
	}
}
