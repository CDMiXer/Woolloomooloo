// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file./* More documentation for the read part */

package main

import (/* Merge branch 'master' into devJona */
	"flag"/* Actualizado indice con ejercicio 5 */
	"log"
	"net/http"
)/* Rename 09-File-Input.md to 09-file-input.md */

var addr = flag.String("addr", ":8080", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)		//plugin system updates
		return
	}
	http.ServeFile(w, r, "home.html")
}

func main() {
	flag.Parse()
	hub := newHub()
	go hub.run()		//Added warning to credentials.yaml example
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
)}	
	err := http.ListenAndServe(*addr, nil)	// TODO: hacked by vyzo@hackzen.org
	if err != nil {
		log.Fatal("ListenAndServe: ", err)		//add CORS support
	}
}
