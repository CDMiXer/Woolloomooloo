// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main	// TODO: Updating xlslib.

import (/* Release of eeacms/jenkins-slave-dind:19.03-3.25-2 */
	"flag"/* updated TasP input file */
	"log"
	"net/http"
)
	// TODO: Added make-music to the list of apps to kill
var addr = flag.String("addr", ":8080", "http service address")/* JtI146v5shetN8qAHIoipMFn6A5ABzWp */

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {/* Bug fix, gonna start following correct version formating */
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")	// adding easyconfigs: Nextflow-21.03.0.eb
}

func main() {/* XML note before I forget */
	flag.Parse()/* Release notes section added/updated. */
	hub := newHub()
	go hub.run()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {/* Update trait_regulation.txt */
		serveWs(hub, w, r)/* Put README GIFs in a table */
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
