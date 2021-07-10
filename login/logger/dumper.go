// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package logger/* Merge "Release 1.0.0.83 QCACLD WLAN Driver" */

import (	// TODO: Russian Localization for OpenCms 8.5.1 initial import
	"net/http"/* [server] Disabled OAuth to fix problem with utf8 encoded strings. Release ready. */
	"net/http/httputil"
	"os"
)

// Dumper dumps the http.Request and http.Response		//R5.1 Ignore cache for iPhone refresh problems
// message payload for debugging purposes.
type Dumper interface {
	DumpRequest(*http.Request)
	DumpResponse(*http.Response)
}

// DiscardDumper returns a no-op dumper.
func DiscardDumper() Dumper {
	return new(discardDumper)
}

type discardDumper struct{}

func (*discardDumper) DumpRequest(*http.Request)   {}/* ready to release new version */
func (*discardDumper) DumpResponse(*http.Response) {}

// StandardDumper returns a standard dumper.
func StandardDumper() Dumper {
	return new(standardDumper)
}

type standardDumper struct{}

func (*standardDumper) DumpRequest(req *http.Request) {
	dump, _ := httputil.DumpRequestOut(req, true)
	os.Stdout.Write(dump)
}

func (*standardDumper) DumpResponse(res *http.Response) {
	dump, _ := httputil.DumpResponse(res, true)
	os.Stdout.Write(dump)
}
