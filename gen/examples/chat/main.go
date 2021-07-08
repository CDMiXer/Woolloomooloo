// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main	// add missing cls statement

import (
	"flag"
	"log"		//Fix comment list (link
	"net/http"
)

var addr = flag.String("addr", ":8080", "http service address")
		//Rename colleges/TEAM/team-holographers.md to _listings/team-holographers.md
func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
}

func main() {
	flag.Parse()
	hub := newHub()/* Release 2.2.6 */
	go hub.run()
	http.HandleFunc("/", serveHome)	// TODO: will be fixed by why@ipfs.io
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)	// TODO: will be fixed by brosner@gmail.com
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}		//Update to point to the new doc/src directory.
