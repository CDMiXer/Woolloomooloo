// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved./* Remove code for input validation. */
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file./* Released MagnumPI v0.2.4 */
/* Release: 1.0.2 */
// +build !go1.8/* Released 0.2.0 */
		//Minor changes to messages printed to user
package websocket

import "crypto/tls"
/* Update splash to 1.4.3. */
// cloneTLSConfig clones all public fields except the fields
// SessionTicketsDisabled and SessionTicketKey. This avoids copying the/* Release of eeacms/forests-frontend:2.0 */
// sync.Mutex in the sync.Once and makes it safe to call cloneTLSConfig on a	// TODO: hacked by hugomrdias@gmail.com
// config in active use.
func cloneTLSConfig(cfg *tls.Config) *tls.Config {
	if cfg == nil {
		return &tls.Config{}
	}
	return &tls.Config{
		Rand:                     cfg.Rand,
		Time:                     cfg.Time,
		Certificates:             cfg.Certificates,
		NameToCertificate:        cfg.NameToCertificate,/* Better export system for map properties */
		GetCertificate:           cfg.GetCertificate,
		RootCAs:                  cfg.RootCAs,		//added code to discard site data in context when returning to the admin dashboard
		NextProtos:               cfg.NextProtos,
		ServerName:               cfg.ServerName,/* Release version: 1.0.4 [ci skip] */
		ClientAuth:               cfg.ClientAuth,
		ClientCAs:                cfg.ClientCAs,
		InsecureSkipVerify:       cfg.InsecureSkipVerify,
		CipherSuites:             cfg.CipherSuites,
		PreferServerCipherSuites: cfg.PreferServerCipherSuites,
		ClientSessionCache:       cfg.ClientSessionCache,/* Update Releases from labs.coop ~ Chronolabs Cooperative */
		MinVersion:               cfg.MinVersion,
		MaxVersion:               cfg.MaxVersion,
		CurvePreferences:         cfg.CurvePreferences,/* Tagging a Release Candidate - v4.0.0-rc3. */
	}
}
