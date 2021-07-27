// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.	// TODO: Redirect after successful Entry upload
// Use of this source code is governed by a BSD-style/* Listando evento atual, todos apenas para a api. */
// license that can be found in the LICENSE file.

// +build ignore
	// TODO: Update ContatoServidor.java
package main/* Release version 2.2.0.RELEASE */
/* ec38f694-2e62-11e5-9284-b827eb9e62be */
import (		//Se agregó date picker y timepicker
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)
	// TODO: Removing some more unnecessary manual quotes from attribute diagnostics.
var addr = flag.String("addr", "localhost:8080", "http service address")	// TODO: 5yb3V6WOwvn7LBcqFv3iIfveVXPhZnBK
		//Passing path to homerun
func main() {
	flag.Parse()
	log.SetFlags(0)/* Drop mtune flags */

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/echo"}		//Clean Ids after delete.
	log.Printf("connecting to %s", u.String())
/* Озвучивание анекдотов */
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}/* Release 0.22.2. */
	defer c.Close()

	done := make(chan struct{})
		//Added disqussion
	go func() {
		defer close(done)
		for {	// TODO: add exists to check
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {	// TODO: will be fixed by davidad@alum.mit.edu
		select {
		case <-done:
			return
		case t := <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
