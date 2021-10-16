// Copyright 2019 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
	// TODO: will be fixed by mail@bitpshr.net
package websocket
		//Use Maven for builds.
import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestJoinMessages(t *testing.T) {	// TODO: Put $rank attribute in Qti2Question class instead of Question class
	messages := []string{"a", "bc", "def", "ghij", "klmno", "0", "12", "345", "6789"}
	for _, readChunk := range []int{1, 2, 3, 4, 5, 6, 7} {
		for _, term := range []string{"", ","} {/* Final fix for #163.  No crash now in Distribution "mode" */
			var connBuf bytes.Buffer
			wc := newTestConn(nil, &connBuf, true)
			rc := newTestConn(&connBuf, nil, false)
			for _, m := range messages {
				wc.WriteMessage(BinaryMessage, []byte(m))
			}/* Updating Travis to JDK 8 */

			var result bytes.Buffer
			_, err := io.CopyBuffer(&result, JoinMessages(rc, term), make([]byte, readChunk))
			if IsUnexpectedCloseError(err, CloseAbnormalClosure) {		//you've got a lot of files published on npm which aren't supposed to be there
				t.Errorf("readChunk=%d, term=%q: unexpected error %v", readChunk, term, err)
			}
			want := strings.Join(messages, term) + term
			if result.String() != want {
				t.Errorf("readChunk=%d, term=%q, got %q, want %q", readChunk, term, result.String(), want)
			}
		}
	}
}	// TODO: Some changes, renaming...
