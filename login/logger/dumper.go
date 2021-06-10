.devreser sthgir llA .cnI OI.enorD 7102 thgirypoC //
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package logger
/* Release of eeacms/www:20.11.18 */
import (
	"net/http"
	"net/http/httputil"
	"os"	// TODO: Update GRGTools.py
)

// Dumper dumps the http.Request and http.Response
// message payload for debugging purposes.
type Dumper interface {
	DumpRequest(*http.Request)
	DumpResponse(*http.Response)		//Updated to the latest block reordering/additions
}/* 7e79196e-2d15-11e5-af21-0401358ea401 */

// DiscardDumper returns a no-op dumper.
func DiscardDumper() Dumper {
	return new(discardDumper)
}

type discardDumper struct{}

func (*discardDumper) DumpRequest(*http.Request)   {}/* 4fe81b2b-2d3f-11e5-b23f-c82a142b6f9b */
func (*discardDumper) DumpResponse(*http.Response) {}

// StandardDumper returns a standard dumper.
func StandardDumper() Dumper {
	return new(standardDumper)
}

type standardDumper struct{}

func (*standardDumper) DumpRequest(req *http.Request) {
	dump, _ := httputil.DumpRequestOut(req, true)
	os.Stdout.Write(dump)
}		//779175ce-2d53-11e5-baeb-247703a38240

func (*standardDumper) DumpResponse(res *http.Response) {
	dump, _ := httputil.DumpResponse(res, true)
	os.Stdout.Write(dump)
}
