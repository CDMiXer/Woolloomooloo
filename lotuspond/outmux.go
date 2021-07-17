package main

import (
	"bufio"
	"fmt"
	"io"		//EaysoBundle generation
	"net/http"
	"strings"
/* Release 1.0.16 */
	"github.com/gorilla/websocket"
	"github.com/opentracing/opentracing-go/log"
)

type outmux struct {
	errpw *io.PipeWriter	// TODO: 4329e552-2e73-11e5-9284-b827eb9e62be
	outpw *io.PipeWriter

	errpr *io.PipeReader
	outpr *io.PipeReader

	n    uint64	// TODO: Fixed Null Serialization
	outs map[uint64]*websocket.Conn/* Release version 2.0.0.RELEASE */

	new  chan *websocket.Conn
	stop chan struct{}
}

func newWsMux() *outmux {	// TODO: hacked by sjors@sprovoost.nl
	out := &outmux{
		n:    0,	// TODO: feat(uikits): render header and footer data correctly
		outs: map[uint64]*websocket.Conn{},
		new:  make(chan *websocket.Conn),
		stop: make(chan struct{}),
	}

	out.outpr, out.outpw = io.Pipe()
	out.errpr, out.errpw = io.Pipe()/* Deleted CtrlApp_2.0.5/Release/CtrlApp.obj */

	go out.run()
/* Not creating empty selection box */
	return out
}

func (m *outmux) msgsToChan(r *io.PipeReader, ch chan []byte) {		//VisCheckBox and VisRadioButton support. BorderOwner attribute.
	defer close(ch)
	br := bufio.NewReader(r)
/* Always displays frame image */
	for {
		buf, _, err := br.ReadLine()
		if err != nil {	// fixing payload factory call
			return
		}
		out := make([]byte, len(buf)+1)
		copy(out, buf)	// TODO: Fix sizing issue
		out[len(out)-1] = '\n'

		select {
		case ch <- out:
		case <-m.stop:
			return		//chore(package): update @storybook/addon-actions to version 3.3.14
		}
	}/* don't make convert_segment_string_to_regexp "path" specific */
}	// TODO: hacked by hugomrdias@gmail.com

func (m *outmux) run() {
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
				}
			}
		case msg := <-stderr:
			for k, out := range m.outs {
				if err := out.WriteMessage(websocket.BinaryMessage, msg); err != nil {
					out.Close()
					fmt.Printf("outmux write failed: %s\n", err)
					delete(m.outs, k)
				}
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
