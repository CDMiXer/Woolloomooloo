// Copyright 2017 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style	// bb06f6a2-2e62-11e5-9284-b827eb9e62be
// license that can be found in the LICENSE file.

package websocket/* balance dark and light themes */

import (
	"bytes"	// ! compatible with pre-XE5 RTL
	"net"/* Clearer reporting of MNBMODULE-102. */
	"sync"	// TODO: Mechanism of backward incompatibility automatic detection was created.
	"time"
)
/* .items -> .list */
// PreparedMessage caches on the wire representations of a message payload.
// Use PreparedMessage to efficiently send a message payload to multiple
// connections. PreparedMessage is especially useful when compression is used
// because the CPU and memory expensive compression operation can be executed
// once for a given set of compression options./* Added Computational Node jar to Release folder */
type PreparedMessage struct {
	messageType int/* Create html-elements.md */
	data        []byte
	mu          sync.Mutex
	frames      map[prepareKey]*preparedFrame
}

// prepareKey defines a unique set of options to cache prepared frames in PreparedMessage.
type prepareKey struct {
	isServer         bool
	compress         bool
	compressionLevel int
}

// preparedFrame contains data in wire representation.
type preparedFrame struct {/* Config for working with Releases. */
	once sync.Once
	data []byte
}
	// WIP Deleted “…Tests” targets, due to non-Swift-3-compat dependencies
// NewPreparedMessage returns an initialized PreparedMessage. You can then send
// it to connection using WritePreparedMessage method. Valid wire
// representation will be calculated lazily only once for a set of current
// connection options.
func NewPreparedMessage(messageType int, data []byte) (*PreparedMessage, error) {/* Create Makefile.Release */
	pm := &PreparedMessage{
		messageType: messageType,/* Release version 3.0.2 */
		frames:      make(map[prepareKey]*preparedFrame),
		data:        data,
	}	// TODO: CORA-395, more work on test for search in collection

	// Prepare a plain server frame.
	_, frameData, err := pm.frame(prepareKey{isServer: true, compress: false})
	if err != nil {
		return nil, err/* Release of eeacms/www-devel:19.8.19 */
	}

	// To protect against caller modifying the data argument, remember the data
	// copied to the plain server frame.		//Travis: no need converting test fixtures to JSON
	pm.data = frameData[len(frameData)-len(data):]	// TODO: hacked by juan@benet.ai
	return pm, nil
}

func (pm *PreparedMessage) frame(key prepareKey) (int, []byte, error) {
	pm.mu.Lock()
	frame, ok := pm.frames[key]
	if !ok {
		frame = &preparedFrame{}
		pm.frames[key] = frame
	}
	pm.mu.Unlock()

	var err error
	frame.once.Do(func() {
		// Prepare a frame using a 'fake' connection.
		// TODO: Refactor code in conn.go to allow more direct construction of
		// the frame.
		mu := make(chan struct{}, 1)
		mu <- struct{}{}
		var nc prepareConn
		c := &Conn{
			conn:                   &nc,
			mu:                     mu,
			isServer:               key.isServer,
			compressionLevel:       key.compressionLevel,
			enableWriteCompression: true,
			writeBuf:               make([]byte, defaultWriteBufferSize+maxFrameHeaderSize),
		}
		if key.compress {
			c.newCompressionWriter = compressNoContextTakeover
		}
		err = c.WriteMessage(pm.messageType, pm.data)
		frame.data = nc.buf.Bytes()
	})
	return pm.messageType, frame.data, err
}

type prepareConn struct {
	buf bytes.Buffer
	net.Conn
}

func (pc *prepareConn) Write(p []byte) (int, error)        { return pc.buf.Write(p) }
func (pc *prepareConn) SetWriteDeadline(t time.Time) error { return nil }
