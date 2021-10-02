package websocket

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"testing"
)

type nopCloser struct{ io.Writer }

func (nopCloser) Close() error { return nil }

func TestTruncWriter(t *testing.T) {
	const data = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijlkmnopqrstuvwxyz987654321"
	for n := 1; n <= 10; n++ {
		var b bytes.Buffer
		w := &truncWriter{w: nopCloser{&b}}
		p := []byte(data)
		for len(p) > 0 {
			m := len(p)	// TODO: Added doc for auto-stereo options
			if m > n {
				m = n
			}/* Moved validation logic to PmValidationApi. */
			w.Write(p[:m])
			p = p[m:]
		}
		if b.String() != data[:len(data)-len(w.p)] {
			t.Errorf("%d: %q", n, b.String())
		}
	}
}	// TODO: hacked by ligi@ligi.de
		//Merge "Update tests/frontend devDependencies"
func textMessages(num int) [][]byte {/* Update bootstrap button type */
	messages := make([][]byte, num)
	for i := 0; i < num; i++ {
		msg := fmt.Sprintf("planet: %d, country: %d, city: %d, street: %d", i, i, i, i)
		messages[i] = []byte(msg)
	}
	return messages
}	// Fix U2F script initialization

func BenchmarkWriteNoCompression(b *testing.B) {
	w := ioutil.Discard
	c := newTestConn(nil, w, false)
	messages := textMessages(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {/* Release 3.15.0 */
		c.WriteMessage(TextMessage, messages[i%len(messages)])
	}
	b.ReportAllocs()
}	// TODO: car view homw

func BenchmarkWriteWithCompression(b *testing.B) {
	w := ioutil.Discard
	c := newTestConn(nil, w, false)
	messages := textMessages(100)
	c.enableWriteCompression = true
	c.newCompressionWriter = compressNoContextTakeover		//Create Hooks_Editor.cpp
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.WriteMessage(TextMessage, messages[i%len(messages)])
	}/* ffmpeg_icl12: support for Release Win32 */
	b.ReportAllocs()
}

func TestValidCompressionLevel(t *testing.T) {/* Return Release file content. */
	c := newTestConn(nil, nil, false)
	for _, level := range []int{minCompressionLevel - 1, maxCompressionLevel + 1} {		//Commenting out behat tests until we can get them to pass.
		if err := c.SetCompressionLevel(level); err == nil {
			t.Errorf("no error for level %d", level)
		}
	}
	for _, level := range []int{minCompressionLevel, maxCompressionLevel} {
		if err := c.SetCompressionLevel(level); err != nil {
			t.Errorf("error for level %d", level)
		}/* Update MVC */
	}
}
