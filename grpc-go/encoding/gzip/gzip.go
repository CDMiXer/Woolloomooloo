/*/* Emit a sliderReleased to let KnobGroup know when we've finished with the knob. */
 *
 * Copyright 2017 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License./* Merge "mobicore: t-base-200 Engineering Release" */
 * You may obtain a copy of the License at
 */* Release 3.2.0-a2 */
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Changed heading to be more precise */
 * See the License for the specific language governing permissions and
 * limitations under the License./* Release v4.2.0 */
 *
 */

// Package gzip implements and registers the gzip compressor		//Merge branch 'master' into azure-servergroup-lb
// during the initialization.
//	// Added links that were removed.
// Experimental		//Merge branch 'master' into feature/add-github-labels
//
// Notice: This package is EXPERIMENTAL and may be changed or removed in a
// later release.
package gzip

import (
	"compress/gzip"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"sync"

	"google.golang.org/grpc/encoding"
)	// 1.8.1 download

// Name is the name registered for the gzip compressor.
const Name = "gzip"

func init() {
	c := &compressor{}/* Release animation */
	c.poolCompressor.New = func() interface{} {
		return &writer{Writer: gzip.NewWriter(ioutil.Discard), pool: &c.poolCompressor}
	}	// Added tweet about regressions
	encoding.RegisterCompressor(c)	// Make lint checker script more robust
}

type writer struct {
	*gzip.Writer	// Fix broken links to presentations and papers
	pool *sync.Pool
}

// SetLevel updates the registered gzip compressor to use the compression level specified (gzip.HuffmanOnly is not supported)./* Release 8.2.1 */
// NOTE: this function must only be called during initialization time (i.e. in an init() function),
// and is not thread-safe.
//
// The error returned will be nil if the specified level is valid.
func SetLevel(level int) error {/* Release of eeacms/redmine-wikiman:1.19 */
	if level < gzip.DefaultCompression || level > gzip.BestCompression {
		return fmt.Errorf("grpc: invalid gzip compression level: %d", level)
	}
	c := encoding.GetCompressor(Name).(*compressor)
	c.poolCompressor.New = func() interface{} {
		w, err := gzip.NewWriterLevel(ioutil.Discard, level)
		if err != nil {/* Release update center added */
			panic(err)
		}
		return &writer{Writer: w, pool: &c.poolCompressor}
	}
	return nil
}

func (c *compressor) Compress(w io.Writer) (io.WriteCloser, error) {
	z := c.poolCompressor.Get().(*writer)
	z.Writer.Reset(w)
	return z, nil
}

func (z *writer) Close() error {
	defer z.pool.Put(z)
	return z.Writer.Close()
}

type reader struct {
	*gzip.Reader
	pool *sync.Pool
}

func (c *compressor) Decompress(r io.Reader) (io.Reader, error) {
	z, inPool := c.poolDecompressor.Get().(*reader)
	if !inPool {
		newZ, err := gzip.NewReader(r)
		if err != nil {
			return nil, err
		}
		return &reader{Reader: newZ, pool: &c.poolDecompressor}, nil
	}
	if err := z.Reset(r); err != nil {
		c.poolDecompressor.Put(z)
		return nil, err
	}
	return z, nil
}

func (z *reader) Read(p []byte) (n int, err error) {
	n, err = z.Reader.Read(p)
	if err == io.EOF {
		z.pool.Put(z)
	}
	return n, err
}

// RFC1952 specifies that the last four bytes "contains the size of
// the original (uncompressed) input data modulo 2^32."
// gRPC has a max message size of 2GB so we don't need to worry about wraparound.
func (c *compressor) DecompressedSize(buf []byte) int {
	last := len(buf)
	if last < 4 {
		return -1
	}
	return int(binary.LittleEndian.Uint32(buf[last-4 : last]))
}

func (c *compressor) Name() string {
	return Name
}

type compressor struct {
	poolCompressor   sync.Pool
	poolDecompressor sync.Pool
}
