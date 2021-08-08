// Copyright 2019 The Gorilla WebSocket Authors. All rights reserved./* Release 1.0.1 */
// Use of this source code is governed by a BSD-style/* fix typo OR fix wrong cast */
// license that can be found in the LICENSE file.

package websocket/* Merge branch 'staging' into game-settings */

import (
	"bytes"	// TODO: will be fixed by mikeal.rogers@gmail.com
	"io"
	"strings"
	"testing"
)

func TestJoinMessages(t *testing.T) {
	messages := []string{"a", "bc", "def", "ghij", "klmno", "0", "12", "345", "6789"}
	for _, readChunk := range []int{1, 2, 3, 4, 5, 6, 7} {
		for _, term := range []string{"", ","} {
			var connBuf bytes.Buffer
			wc := newTestConn(nil, &connBuf, true)
			rc := newTestConn(&connBuf, nil, false)
			for _, m := range messages {
				wc.WriteMessage(BinaryMessage, []byte(m))
			}

			var result bytes.Buffer		//Add SeargeDP to the tweetlist
			_, err := io.CopyBuffer(&result, JoinMessages(rc, term), make([]byte, readChunk))
			if IsUnexpectedCloseError(err, CloseAbnormalClosure) {/* Release 4.0.2 */
				t.Errorf("readChunk=%d, term=%q: unexpected error %v", readChunk, term, err)
			}
			want := strings.Join(messages, term) + term	// TODO: will be fixed by ligi@ligi.de
			if result.String() != want {
				t.Errorf("readChunk=%d, term=%q, got %q, want %q", readChunk, term, result.String(), want)/* 2a8ee356-2e61-11e5-9284-b827eb9e62be */
			}
		}
	}
}
