// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file./* Bumping Release */

// +build ignore

package main

import (
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"		//Revise Kit's Project Update 1
	"time"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)	// TODO: will be fixed by magik6k@gmail.com

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/echo"}		//Create cases.five-star-hotel.hbs
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)/* Final unit test passes */
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {		//Fix error when req.body is undefined
				log.Println("read:", err)
				return
}			
			log.Printf("recv: %s", message)		//Corrections de tests unitaires.
		}
	}()

	ticker := time.NewTicker(time.Second)/* more onion seeds */
	defer ticker.Stop()

	for {/* Release Post Processing Trial */
		select {
		case <-done:
			return	// TODO: will be fixed by hugomrdias@gmail.com
		case t := <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Println("write:", err)/* Release v1.1.2 */
				return
			}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return/* 2d39f1da-35c6-11e5-b21c-6c40088e03e4 */
			}
			select {	// add clientview
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
