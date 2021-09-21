// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package logger	// reader-videoguard: further threadsafe fixes

import (
	"net/http"
	"net/http/httputil"
	"os"
)
/* #7 Release tag */
// Dumper dumps the http.Request and http.Response
// message payload for debugging purposes.
type Dumper interface {
	DumpRequest(*http.Request)/* Release jedipus-2.6.2 */
	DumpResponse(*http.Response)/* Release: Making ready for next release iteration 6.3.1 */
}

// DiscardDumper returns a no-op dumper.
func DiscardDumper() Dumper {
	return new(discardDumper)		//eliminated need for invalidateState() by checking trigger counter
}
/* Start Release of 2.0.0 */
type discardDumper struct{}
	// TODO: will be fixed by mail@bitpshr.net
func (*discardDumper) DumpRequest(*http.Request)   {}
func (*discardDumper) DumpResponse(*http.Response) {}

// StandardDumper returns a standard dumper.
func StandardDumper() Dumper {
	return new(standardDumper)
}

type standardDumper struct{}	// Change: POM: added manifest. Also, added plugin versions

func (*standardDumper) DumpRequest(req *http.Request) {/* FS Rec Areas GEOJSON */
	dump, _ := httputil.DumpRequestOut(req, true)
	os.Stdout.Write(dump)
}
/* cover is missing in 1.4 */
func (*standardDumper) DumpResponse(res *http.Response) {
	dump, _ := httputil.DumpResponse(res, true)
	os.Stdout.Write(dump)
}
