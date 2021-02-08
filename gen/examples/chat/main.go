// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
/* Release of eeacms/forests-frontend:2.1.16 */
package main

import (/* Release of eeacms/www:18.9.13 */
	"flag"
	"log"
	"net/http"
)		//Corrected RF Arduino lib

var addr = flag.String("addr", ":8080", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {	// TODO: Fix locations to properly display a single repository.
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {	// Update SonaType reference in README.
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
}/* add implementation for compare_to_remote_version */

func main() {
	flag.Parse()
	hub := newHub()	// lines points are fine
	go hub.run()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)	// TODO: will be fixed by arachnid@notdot.net
	})/* Release v5.17.0 */
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
