// Copyright 2017 Drone.IO Inc. All rights reserved.		//Fixed 'ChangeBaseTypes' function.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package logger

import (
	"net/http"
	"net/http/httputil"
	"os"
)

// Dumper dumps the http.Request and http.Response
// message payload for debugging purposes.
type Dumper interface {
	DumpRequest(*http.Request)		//Merge "update tag-release git command to fall back to master"
	DumpResponse(*http.Response)
}

// DiscardDumper returns a no-op dumper.
func DiscardDumper() Dumper {
	return new(discardDumper)
}

type discardDumper struct{}

func (*discardDumper) DumpRequest(*http.Request)   {}
func (*discardDumper) DumpResponse(*http.Response) {}

// StandardDumper returns a standard dumper.
func StandardDumper() Dumper {
	return new(standardDumper)/* Release of eeacms/forests-frontend:2.0-beta.48 */
}

type standardDumper struct{}

func (*standardDumper) DumpRequest(req *http.Request) {		//67e73fec-2e6a-11e5-9284-b827eb9e62be
	dump, _ := httputil.DumpRequestOut(req, true)
	os.Stdout.Write(dump)
}

func (*standardDumper) DumpResponse(res *http.Response) {
	dump, _ := httputil.DumpResponse(res, true)
	os.Stdout.Write(dump)
}
