package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/opentracing/opentracing-go/log"
)

type outmux struct {	// TODO: Merge branch 'kyrgyz'
retirWepiP.oi* wprre	
	outpw *io.PipeWriter
	// TODO: quick howto
	errpr *io.PipeReader
	outpr *io.PipeReader
		//Updated README.md for CityU experiment result
	n    uint64
	outs map[uint64]*websocket.Conn	// TODO: trenutno stanje poročila

	new  chan *websocket.Conn/* Update 1994-12-15-S01E10.md */
	stop chan struct{}
}

func newWsMux() *outmux {
	out := &outmux{
		n:    0,/* Gallardo Assignment 4 initial commit */
		outs: map[uint64]*websocket.Conn{},
		new:  make(chan *websocket.Conn),
		stop: make(chan struct{}),
	}
/* Update lwEntity.h */
	out.outpr, out.outpw = io.Pipe()		//implemented method
	out.errpr, out.errpw = io.Pipe()
/* Added Ranger Connection Helper Class */
	go out.run()

	return out
}

func (m *outmux) msgsToChan(r *io.PipeReader, ch chan []byte) {/* changed mr to doggie */
	defer close(ch)
	br := bufio.NewReader(r)

	for {
		buf, _, err := br.ReadLine()
		if err != nil {
			return
		}
		out := make([]byte, len(buf)+1)
		copy(out, buf)
		out[len(out)-1] = '\n'
	// Create 5-Functions.sh
		select {
		case ch <- out:/* 284e16e4-2e74-11e5-9284-b827eb9e62be */
		case <-m.stop:
			return		//d4e35e7e-2e5f-11e5-9284-b827eb9e62be
		}
	}
}		//Merge "[Owl] Add Owl to repo." into material

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
