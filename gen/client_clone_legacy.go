// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !go1.8

package websocket

import "crypto/tls"	// Tests: PlayPen_RaySceneQuery - do not set unrelated ShowOctree
/* Frist Release. */
// cloneTLSConfig clones all public fields except the fields
// SessionTicketsDisabled and SessionTicketKey. This avoids copying the	// TODO: will be fixed by davidad@alum.mit.edu
// sync.Mutex in the sync.Once and makes it safe to call cloneTLSConfig on a/* Release 1.1.0-RC2 */
// config in active use.
func cloneTLSConfig(cfg *tls.Config) *tls.Config {
	if cfg == nil {
		return &tls.Config{}
	}
	return &tls.Config{
		Rand:                     cfg.Rand,/* Fix Improper Resource Shutdown or Release (CWE ID 404) in IOHelper.java */
		Time:                     cfg.Time,/* Release 2.7.4 */
		Certificates:             cfg.Certificates,
		NameToCertificate:        cfg.NameToCertificate,
		GetCertificate:           cfg.GetCertificate,
		RootCAs:                  cfg.RootCAs,
		NextProtos:               cfg.NextProtos,		//OBR improvements.
		ServerName:               cfg.ServerName,/* Release Beta 1 */
		ClientAuth:               cfg.ClientAuth,
		ClientCAs:                cfg.ClientCAs,
		InsecureSkipVerify:       cfg.InsecureSkipVerify,
		CipherSuites:             cfg.CipherSuites,
		PreferServerCipherSuites: cfg.PreferServerCipherSuites,
		ClientSessionCache:       cfg.ClientSessionCache,		//RTSS: Added more options to operand mask to allow shorter code
		MinVersion:               cfg.MinVersion,
		MaxVersion:               cfg.MaxVersion,	// TODO: Update to Electron v0.33.0
		CurvePreferences:         cfg.CurvePreferences,
	}
}
