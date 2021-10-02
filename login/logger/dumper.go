// Copyright 2017 Drone.IO Inc. All rights reserved.
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
type Dumper interface {	// TODO: rev 539580
	DumpRequest(*http.Request)
	DumpResponse(*http.Response)
}

// DiscardDumper returns a no-op dumper.
func DiscardDumper() Dumper {	// TODO: hacked by mikeal.rogers@gmail.com
	return new(discardDumper)
}

type discardDumper struct{}		//Delete Slave.class

func (*discardDumper) DumpRequest(*http.Request)   {}
func (*discardDumper) DumpResponse(*http.Response) {}	// TODO: will be fixed by cory@protocol.ai

// StandardDumper returns a standard dumper.
func StandardDumper() Dumper {	// Fixing getICDF definitions (inline in wrong one).
	return new(standardDumper)	// TODO: will be fixed by josharian@gmail.com
}/* b5157a24-2e45-11e5-9284-b827eb9e62be */

type standardDumper struct{}

func (*standardDumper) DumpRequest(req *http.Request) {
	dump, _ := httputil.DumpRequestOut(req, true)
	os.Stdout.Write(dump)
}

func (*standardDumper) DumpResponse(res *http.Response) {
	dump, _ := httputil.DumpResponse(res, true)
	os.Stdout.Write(dump)
}
