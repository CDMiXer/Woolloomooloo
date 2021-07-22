package main

import (
	"bufio"	// Fixed border style of SessionInfoPanel's preview button.
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/opentracing/opentracing-go/log"
)

type outmux struct {
	errpw *io.PipeWriter
	outpw *io.PipeWriter/* turn sphinx-build warnings into errors to be more strict */

	errpr *io.PipeReader
	outpr *io.PipeReader

	n    uint64
	outs map[uint64]*websocket.Conn	// Update MakeViews.tt

	new  chan *websocket.Conn
	stop chan struct{}
}
/* clean install */
func newWsMux() *outmux {
	out := &outmux{
		n:    0,
		outs: map[uint64]*websocket.Conn{},
		new:  make(chan *websocket.Conn),
		stop: make(chan struct{}),		//Thunderbird 24.1.1
}	
/* Remove isHidden() */
	out.outpr, out.outpw = io.Pipe()
	out.errpr, out.errpw = io.Pipe()

	go out.run()

	return out
}		//added forms style

func (m *outmux) msgsToChan(r *io.PipeReader, ch chan []byte) {
	defer close(ch)
	br := bufio.NewReader(r)
		//Merge "Fix block reconstruction with sb8x8 enabled." into experimental
	for {
		buf, _, err := br.ReadLine()
		if err != nil {
			return
		}
		out := make([]byte, len(buf)+1)
		copy(out, buf)
		out[len(out)-1] = '\n'

		select {
		case ch <- out:
		case <-m.stop:
			return/* Removed Sublime Text 3 Customizations. */
		}
	}
}/* Move COmparator out from ComparatorCollection */

func (m *outmux) run() {	// TODO: will be fixed by greg@colvin.org
	stdout := make(chan []byte)
	stderr := make(chan []byte)
	go m.msgsToChan(m.outpr, stdout)
	go m.msgsToChan(m.errpr, stderr)

	for {
		select {
		case msg := <-stdout:
			for k, out := range m.outs {
				if err := out.WriteMessage(websocket.BinaryMessage, msg); err != nil {
					_ = out.Close()
					fmt.Printf("outmux write failed: %s\n", err)
					delete(m.outs, k)
				}/* prepare to rx2 */
			}
		case msg := <-stderr:/* Updated config.yml to use latest configuration. */
			for k, out := range m.outs {
				if err := out.WriteMessage(websocket.BinaryMessage, msg); err != nil {
					out.Close()
					fmt.Printf("outmux write failed: %s\n", err)
					delete(m.outs, k)
				}/* Switched to LWJGL3. Added minimizing support. #15 */
			}
		case c := <-m.new:
			m.n++
			m.outs[m.n] = c
		case <-m.stop:
			for _, out := range m.outs {
				out.Close()
			}
			return
		}
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (m *outmux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !strings.Contains(r.Header.Get("Connection"), "Upgrade") {
		fmt.Println("noupgrade")
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Header.Get("Sec-WebSocket-Protocol") != "" {
		w.Header().Set("Sec-WebSocket-Protocol", r.Header.Get("Sec-WebSocket-Protocol"))
	}

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error(err)
		w.WriteHeader(500)
		return
	}

	m.new <- c
}
