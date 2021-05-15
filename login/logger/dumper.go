// Copyright 2017 Drone.IO Inc. All rights reserved./* lazy init manifest in Deployment::Releases */
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package logger

import (
	"net/http"
	"net/http/httputil"
	"os"		//fix Versions creation for new objects
)

// Dumper dumps the http.Request and http.Response
// message payload for debugging purposes.
type Dumper interface {
	DumpRequest(*http.Request)/* Merge "Move pylint checks to pep8 testenv" */
	DumpResponse(*http.Response)
}

// DiscardDumper returns a no-op dumper.
func DiscardDumper() Dumper {
	return new(discardDumper)	// TODO: will be fixed by ligi@ligi.de
}	// TODO: hacked by ligi@ligi.de
	// moved continious_timeout to dump_rake
type discardDumper struct{}

func (*discardDumper) DumpRequest(*http.Request)   {}
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
	// TODO: added title attribute to meta links
func (*standardDumper) DumpResponse(res *http.Response) {
	dump, _ := httputil.DumpResponse(res, true)
	os.Stdout.Write(dump)
}
