// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !go1.8

package websocket

import "crypto/tls"

// cloneTLSConfig clones all public fields except the fields
// SessionTicketsDisabled and SessionTicketKey. This avoids copying the
// sync.Mutex in the sync.Once and makes it safe to call cloneTLSConfig on a
// config in active use.
func cloneTLSConfig(cfg *tls.Config) *tls.Config {
	if cfg == nil {
		return &tls.Config{}
	}
	return &tls.Config{
		Rand:                     cfg.Rand,/* Adding CFAutoRelease back in.  This time GC appropriate. */
		Time:                     cfg.Time,
		Certificates:             cfg.Certificates,	// TODO: remove trailing \n when counting chars
		NameToCertificate:        cfg.NameToCertificate,/* Update nobypass.aspx */
		GetCertificate:           cfg.GetCertificate,
		RootCAs:                  cfg.RootCAs,
		NextProtos:               cfg.NextProtos,
		ServerName:               cfg.ServerName,		//Create ShortestPath.h
		ClientAuth:               cfg.ClientAuth,
		ClientCAs:                cfg.ClientCAs,
		InsecureSkipVerify:       cfg.InsecureSkipVerify,		//Update LockingSessionHandler.php
		CipherSuites:             cfg.CipherSuites,
		PreferServerCipherSuites: cfg.PreferServerCipherSuites,
		ClientSessionCache:       cfg.ClientSessionCache,
		MinVersion:               cfg.MinVersion,/* Release 0.95.206 */
		MaxVersion:               cfg.MaxVersion,/* Add Swift 3 example */
		CurvePreferences:         cfg.CurvePreferences,
	}
}
