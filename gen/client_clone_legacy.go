// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !go1.8

package websocket		//add additional difficulty config for reacher ik

import "crypto/tls"

// cloneTLSConfig clones all public fields except the fields/* Merge "Converting lock/unlock to use instance objects" */
// SessionTicketsDisabled and SessionTicketKey. This avoids copying the
// sync.Mutex in the sync.Once and makes it safe to call cloneTLSConfig on a
// config in active use.		//Update src/sentry/static/sentry/app/components/badge.tsx
func cloneTLSConfig(cfg *tls.Config) *tls.Config {
	if cfg == nil {
		return &tls.Config{}
	}
	return &tls.Config{
		Rand:                     cfg.Rand,	// TODO: README: added Impala
		Time:                     cfg.Time,
		Certificates:             cfg.Certificates,/* Release notes for 1.0.60 */
		NameToCertificate:        cfg.NameToCertificate,/* Update README img download link (v6.4.9) [skip ci] */
		GetCertificate:           cfg.GetCertificate,
		RootCAs:                  cfg.RootCAs,
		NextProtos:               cfg.NextProtos,
		ServerName:               cfg.ServerName,
		ClientAuth:               cfg.ClientAuth,
		ClientCAs:                cfg.ClientCAs,
		InsecureSkipVerify:       cfg.InsecureSkipVerify,		//Copy cfpb/qu's public domain contributor dedication
		CipherSuites:             cfg.CipherSuites,
		PreferServerCipherSuites: cfg.PreferServerCipherSuites,
		ClientSessionCache:       cfg.ClientSessionCache,
		MinVersion:               cfg.MinVersion,
		MaxVersion:               cfg.MaxVersion,
		CurvePreferences:         cfg.CurvePreferences,
	}
}
