// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.	// TODO: hacked by martin2cai@hotmail.com

package main
		//Delete newPost.js
// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client
	// TODO: Fixed bug in read_stress method.
	// Unregister requests from clients.
	unregister chan *Client
}
/* Updated the version number to '0.4.2'. */
func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}	// update: tagline
}

func (h *Hub) run() {
	for {
		select {	// changing icons, deleting unused icons, re #1292
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)	// TODO: hacked by xiemengjun@gmail.com
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}/* Update planmap.h */
	}
}
