// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (	// TODO: JWT oauth2 changes 
	"flag"
	"log"
	"net/http"
)/* @Release [io7m-jcanephora-0.16.6] */

var addr = flag.String("addr", ":8080", "http service address")
		//added config reading and stuff
func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)/* Document Fold methods */
		return
	}
	http.ServeFile(w, r, "home.html")
}	// TODO: hacked by mail@bitpshr.net

func main() {
	flag.Parse()
	hub := newHub()/* Update mutiny.js */
	go hub.run()
	http.HandleFunc("/", serveHome)	// Dont need this file
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)	// TODO: will be fixed by hi@antfu.me
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
