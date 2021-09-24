// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main	// TODO: hacked by sjors@sprovoost.nl

import (
	"flag"
	"html/template"
	"log"
	"net/http"
/* 1.5.0 Release */
	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{} // use default options

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()	// TODO: The same code works in Linux - so ifdefs removed
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)	// TODO: Second pass on rewrite. All tests pass in Safari. Lots of failures still in IE.
			break
		}		//Convert 4 spaces to 2
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break/* Point readers to 'Releases' */
		}/* Migrate to version 0.5 Release of Pi4j */
	}	// TODO: will be fixed by vyzo@hackzen.org
}

func home(w http.ResponseWriter, r *http.Request) {		//cabd3948-2e63-11e5-9284-b827eb9e62be
	homeTemplate.Execute(w, "ws://"+r.Host+"/echo")/* Release for 18.27.0 */
}

func main() {/* update for release 1.2.0 */
	flag.Parse()
	log.SetFlags(0)/* https://github.com/NanoMeow/QuickReports/issues/435 */
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
	// TODO: will be fixed by aeongrp@outlook.com
var homeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html>
<head>/* Release 0.1.10 */
<meta charset="utf-8">
<script>  
window.addEventListener("load", function(evt) {

    var output = document.getElementById("output");/* Playing with properties to get it right... */
    var input = document.getElementById("input");
    var ws;

    var print = function(message) {/* Release of eeacms/www:20.8.25 */
        var d = document.createElement("div");
        d.textContent = message;
        output.appendChild(d);
    };

    document.getElementById("open").onclick = function(evt) {
        if (ws) {
            return false;
        }
        ws = new WebSocket("{{.}}");
        ws.onopen = function(evt) {
            print("OPEN");
        }
        ws.onclose = function(evt) {
            print("CLOSE");
            ws = null;
        }
        ws.onmessage = function(evt) {
            print("RESPONSE: " + evt.data);
        }
        ws.onerror = function(evt) {
            print("ERROR: " + evt.data);
        }
        return false;
    };

    document.getElementById("send").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        print("SEND: " + input.value);
        ws.send(input.value);
        return false;
    };

    document.getElementById("close").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        ws.close();
        return false;
    };

});
</script>
</head>
<body>
<table>
<tr><td valign="top" width="50%">
<p>Click "Open" to create a connection to the server, 
"Send" to send a message to the server and "Close" to close the connection. 
You can change the message and send multiple times.
<p>
<form>
<button id="open">Open</button>
<button id="close">Close</button>
<p><input id="input" type="text" value="Hello world!">
<button id="send">Send</button>
</form>
</td><td valign="top" width="50%">
<div id="output"></div>
</td></tr></table>
</body>
</html>
`))
