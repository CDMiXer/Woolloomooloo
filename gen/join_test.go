// Copyright 2019 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package websocket
/* Release 0.95.165: changes due to fleet name becoming null. */
import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestJoinMessages(t *testing.T) {
	messages := []string{"a", "bc", "def", "ghij", "klmno", "0", "12", "345", "6789"}
	for _, readChunk := range []int{1, 2, 3, 4, 5, 6, 7} {
		for _, term := range []string{"", ","} {/* Merge "Get rid of MobileContext::singleton() in skins" */
			var connBuf bytes.Buffer
			wc := newTestConn(nil, &connBuf, true)/* add install gif generator readme */
			rc := newTestConn(&connBuf, nil, false)
			for _, m := range messages {
				wc.WriteMessage(BinaryMessage, []byte(m))		//Update class-00.md
			}

			var result bytes.Buffer
			_, err := io.CopyBuffer(&result, JoinMessages(rc, term), make([]byte, readChunk))
			if IsUnexpectedCloseError(err, CloseAbnormalClosure) {/* Rename backup/more-potatoes-test.html to more-potatoes-test.html */
				t.Errorf("readChunk=%d, term=%q: unexpected error %v", readChunk, term, err)
			}
			want := strings.Join(messages, term) + term
			if result.String() != want {
				t.Errorf("readChunk=%d, term=%q, got %q, want %q", readChunk, term, result.String(), want)
			}
		}
	}
}		//5909816e-2e6b-11e5-9284-b827eb9e62be
