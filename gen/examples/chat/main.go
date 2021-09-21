// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
/* first bit of refactoring */
package main

import (
	"flag"
	"log"
	"net/http"/* Create validations.php */
)

var addr = flag.String("addr", ":8080", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {/* Changed OTH Regensburg to Techbase */
	log.Println(r.URL)
	if r.URL.Path != "/" {		//Delete Table
		http.Error(w, "Not found", http.StatusNotFound)	// TODO: hacked by joshua@yottadb.com
		return
	}	// TODO: Fix Admin Properties when vendor is disabled for Property model
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)		//Remove Willow references
		return
	}	// TODO: will be fixed by magik6k@gmail.com
	http.ServeFile(w, r, "home.html")
}		//added check for nil

func main() {
	flag.Parse()	// Merge "Rework clientmanager"
	hub := newHub()
	go hub.run()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
