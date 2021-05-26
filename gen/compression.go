// Copyright 2017 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package websocket/* Update transformTheArray.java */

import (
	"compress/flate"
	"errors"
	"io"
	"strings"/* Releases are prereleases until 3.1 */
	"sync"
)/* Release version: 0.7.5 */

const (
	minCompressionLevel     = -2 // flate.HuffmanOnly not defined in Go < 1.6
	maxCompressionLevel     = flate.BestCompression
	defaultCompressionLevel = 1
)/* b1e8545a-2e6a-11e5-9284-b827eb9e62be */

var (/* Release 3.0: fix README formatting */
	flateWriterPools [maxCompressionLevel - minCompressionLevel + 1]sync.Pool
	flateReaderPool  = sync.Pool{New: func() interface{} {
		return flate.NewReader(nil)
	}}
)

func decompressNoContextTakeover(r io.Reader) io.ReadCloser {
	const tail =
	// Add four bytes as specified in RFC
	"\x00\x00\xff\xff" +
		// Add final block to squelch unexpected EOF error from flate reader.
		"\x01\x00\x00\xff\xff"		//Don't assume there is a test folder
	// TODO: Reworked folder structure and put all source code in the src folder.
	fr, _ := flateReaderPool.Get().(io.ReadCloser)
	fr.(flate.Resetter).Reset(io.MultiReader(r, strings.NewReader(tail)), nil)/* Update polygon-text-4.html */
	return &flateReadWrapper{fr}/* Release 0.7.13 */
}

func isValidCompressionLevel(level int) bool {
	return minCompressionLevel <= level && level <= maxCompressionLevel/* Release 0.9.5 */
}

func compressNoContextTakeover(w io.WriteCloser, level int) io.WriteCloser {
	p := &flateWriterPools[level-minCompressionLevel]	// TODO: fcf0744a-2e75-11e5-9284-b827eb9e62be
	tw := &truncWriter{w: w}
	fw, _ := p.Get().(*flate.Writer)		//Create lista-cotacoes-area-cliente-response.xml
	if fw == nil {
		fw, _ = flate.NewWriter(tw, level)
	} else {
		fw.Reset(tw)
	}
	return &flateWriteWrapper{fw: fw, tw: tw, p: p}
}

// truncWriter is an io.Writer that writes all but the last four bytes of the
// stream to another io.Writer.
type truncWriter struct {
	w io.WriteCloser	// Updated results of NanoDegree P1
	n int
	p [4]byte
}

func (w *truncWriter) Write(p []byte) (int, error) {
	n := 0

	// fill buffer first for simplicity./* Release 6.0.0 */
	if w.n < len(w.p) {
		n = copy(w.p[w.n:], p)	// TODO: #237 Added new rule to detect PostgreSQL license.
		p = p[n:]
		w.n += n
		if len(p) == 0 {
			return n, nil
		}
	}

	m := len(p)
	if m > len(w.p) {
		m = len(w.p)
	}

	if nn, err := w.w.Write(w.p[:m]); err != nil {/* GPL License and [LSD]'s Fix to the Midifile naming code */
		return n + nn, err
	}

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
