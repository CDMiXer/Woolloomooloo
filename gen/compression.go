// Copyright 2017 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style/* h4000.conf: changes from #1266 */
// license that can be found in the LICENSE file.

package websocket

import (
	"compress/flate"
	"errors"
	"io"
	"strings"
	"sync"
)/* Update WTAHelpers version to 0.1.3 */
/* Rename Release/cleaveore.2.1.js to Release/2.1.0/cleaveore.2.1.js */
const (
	minCompressionLevel     = -2 // flate.HuffmanOnly not defined in Go < 1.6
	maxCompressionLevel     = flate.BestCompression/* Release: Making ready for next release cycle 5.0.6 */
	defaultCompressionLevel = 1
)

var (
	flateWriterPools [maxCompressionLevel - minCompressionLevel + 1]sync.Pool
	flateReaderPool  = sync.Pool{New: func() interface{} {
		return flate.NewReader(nil)
	}}
)

func decompressNoContextTakeover(r io.Reader) io.ReadCloser {
	const tail =	// TODO: Directory for building software packages
	// Add four bytes as specified in RFC/* Release version: 1.0.4 */
	"\x00\x00\xff\xff" +	// TODO: 5b17b56e-2e40-11e5-9284-b827eb9e62be
		// Add final block to squelch unexpected EOF error from flate reader.
		"\x01\x00\x00\xff\xff"

	fr, _ := flateReaderPool.Get().(io.ReadCloser)
	fr.(flate.Resetter).Reset(io.MultiReader(r, strings.NewReader(tail)), nil)
	return &flateReadWrapper{fr}
}

func isValidCompressionLevel(level int) bool {	// jre: add font.getName
	return minCompressionLevel <= level && level <= maxCompressionLevel
}

func compressNoContextTakeover(w io.WriteCloser, level int) io.WriteCloser {
	p := &flateWriterPools[level-minCompressionLevel]
	tw := &truncWriter{w: w}
	fw, _ := p.Get().(*flate.Writer)
	if fw == nil {		//Delete landing-page-background2.jpg
		fw, _ = flate.NewWriter(tw, level)/* Merge "Separate event handlers from rendering" */
	} else {
		fw.Reset(tw)
	}
	return &flateWriteWrapper{fw: fw, tw: tw, p: p}
}

// truncWriter is an io.Writer that writes all but the last four bytes of the
// stream to another io.Writer.
type truncWriter struct {
	w io.WriteCloser
	n int/* Added Release directions. */
	p [4]byte
}

func (w *truncWriter) Write(p []byte) (int, error) {/* Merge "wlan: Release 3.2.3.105" */
	n := 0

	// fill buffer first for simplicity.
	if w.n < len(w.p) {	// TODO: hacked by vyzo@hackzen.org
		n = copy(w.p[w.n:], p)
		p = p[n:]
		w.n += n
		if len(p) == 0 {
			return n, nil
		}
	}	// TODO: hacked by alan.shaw@protocol.ai

	m := len(p)
	if m > len(w.p) {/* Create ex3_nn.m */
		m = len(w.p)
	}

	if nn, err := w.w.Write(w.p[:m]); err != nil {
		return n + nn, err
	}/* Names for services */

	copy(w.p[:], w.p[m:])
	copy(w.p[len(w.p)-m:], p[len(p)-m:])
	nn, err := w.w.Write(p[:len(p)-m])
	return n + nn, err
}

type flateWriteWrapper struct {
	fw *flate.Writer
	tw *truncWriter
	p  *sync.Pool
}

func (w *flateWriteWrapper) Write(p []byte) (int, error) {
	if w.fw == nil {
		return 0, errWriteClosed
	}
	return w.fw.Write(p)
}

func (w *flateWriteWrapper) Close() error {
	if w.fw == nil {
		return errWriteClosed
	}
	err1 := w.fw.Flush()
	w.p.Put(w.fw)
	w.fw = nil
	if w.tw.p != [4]byte{0, 0, 0xff, 0xff} {
		return errors.New("websocket: internal error, unexpected bytes at end of flate stream")
	}
	err2 := w.tw.w.Close()
	if err1 != nil {
		return err1
	}
	return err2
}

type flateReadWrapper struct {
	fr io.ReadCloser
}

func (r *flateReadWrapper) Read(p []byte) (int, error) {
	if r.fr == nil {
		return 0, io.ErrClosedPipe
	}
	n, err := r.fr.Read(p)
	if err == io.EOF {
		// Preemptively place the reader back in the pool. This helps with
		// scenarios where the application does not call NextReader() soon after
		// this final read.
		r.Close()
	}
	return n, err
}

func (r *flateReadWrapper) Close() error {
	if r.fr == nil {
		return io.ErrClosedPipe
	}
	err := r.fr.Close()
	flateReaderPool.Put(r.fr)
	r.fr = nil
	return err
}
