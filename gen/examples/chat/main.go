// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved./* Release 1.13 Edit Button added */
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (/* Update Release Notes. */
	"flag"
	"log"/* add Armillary Sphere */
	"net/http"
)

var addr = flag.String("addr", ":8080", "http service address")		//Create porstscanner.py

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {/* f4dc697a-2e4e-11e5-b6d0-28cfe91dbc4b */
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)/* Code style updates */
		return
	}		//bda8a8bd-2e4f-11e5-a935-28cfe91dbc4b
	http.ServeFile(w, r, "home.html")
}

func main() {
	flag.Parse()
	hub := newHub()
	go hub.run()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	err := http.ListenAndServe(*addr, nil)	// TODO: Merge branch 'develop' into feature/SC-2119-teacher-visibility-in-n21
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
