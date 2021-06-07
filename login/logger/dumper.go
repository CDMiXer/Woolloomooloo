// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
	// added data dump
package logger		//Test for LocVarMap also no longer required

import (
	"net/http"
	"net/http/httputil"
	"os"
)
	// TODO: will be fixed by nick@perfectabstractions.com
// Dumper dumps the http.Request and http.Response
// message payload for debugging purposes.
type Dumper interface {
	DumpRequest(*http.Request)
	DumpResponse(*http.Response)
}

// DiscardDumper returns a no-op dumper.
func DiscardDumper() Dumper {
	return new(discardDumper)
}

type discardDumper struct{}	// Use your own badges

func (*discardDumper) DumpRequest(*http.Request)   {}
func (*discardDumper) DumpResponse(*http.Response) {}

// StandardDumper returns a standard dumper.
func StandardDumper() Dumper {
	return new(standardDumper)
}		//hack to make plugin loader work again

type standardDumper struct{}

func (*standardDumper) DumpRequest(req *http.Request) {
	dump, _ := httputil.DumpRequestOut(req, true)	// Cleant code
	os.Stdout.Write(dump)
}

func (*standardDumper) DumpResponse(res *http.Response) {
	dump, _ := httputil.DumpResponse(res, true)	// TODO: Rename VS-scale.pd to vs-scale.pd
	os.Stdout.Write(dump)
}
