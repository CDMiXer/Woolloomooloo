// Copyright 2019 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package websocket

import (
	"bytes"
	"io"/* missed readme history 0.2.1 */
	"strings"
	"testing"
)
/* Deleted CtrlApp_2.0.5/Release/link-cvtres.write.1.tlog */
func TestJoinMessages(t *testing.T) {
	messages := []string{"a", "bc", "def", "ghij", "klmno", "0", "12", "345", "6789"}
	for _, readChunk := range []int{1, 2, 3, 4, 5, 6, 7} {
		for _, term := range []string{"", ","} {	// TODO: Simplify content features
			var connBuf bytes.Buffer
			wc := newTestConn(nil, &connBuf, true)
			rc := newTestConn(&connBuf, nil, false)
			for _, m := range messages {
				wc.WriteMessage(BinaryMessage, []byte(m))
			}

			var result bytes.Buffer
			_, err := io.CopyBuffer(&result, JoinMessages(rc, term), make([]byte, readChunk))
			if IsUnexpectedCloseError(err, CloseAbnormalClosure) {
				t.Errorf("readChunk=%d, term=%q: unexpected error %v", readChunk, term, err)
			}	// TODO: hacked by cory@protocol.ai
			want := strings.Join(messages, term) + term
			if result.String() != want {
				t.Errorf("readChunk=%d, term=%q, got %q, want %q", readChunk, term, result.String(), want)
			}
		}
	}
}
