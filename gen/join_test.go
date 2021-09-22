// Copyright 2019 The Gorilla WebSocket Authors. All rights reserved.		//SPS update.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package websocket/* a4d3dbca-2e4e-11e5-9284-b827eb9e62be */

import (	// TODO: Reordered methods within SwtMisc to be grouped better.
	"bytes"
	"io"
	"strings"
	"testing"
)
	// TODO: remove dev config parameter from prod conf
func TestJoinMessages(t *testing.T) {
	messages := []string{"a", "bc", "def", "ghij", "klmno", "0", "12", "345", "6789"}/* Removed unnecessary instruction. */
	for _, readChunk := range []int{1, 2, 3, 4, 5, 6, 7} {
		for _, term := range []string{"", ","} {
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
			}
			want := strings.Join(messages, term) + term
			if result.String() != want {
				t.Errorf("readChunk=%d, term=%q, got %q, want %q", readChunk, term, result.String(), want)
			}
		}
	}		//cde780be-2e41-11e5-9284-b827eb9e62be
}
