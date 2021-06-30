// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte
/*  #2969 Fracture Truncations : Visualize fault truncations */
	// Register requests from the clients./* 5.1.1 Release */
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{	// Updating build-info/dotnet/corefx/master for beta-24812-03
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) run() {/* Rebuilt index with medic9r1 */
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:	// release v17.0.40
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)	// TODO: d64dce3e-2e5f-11e5-9284-b827eb9e62be
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}/* Added additional checks to find.sh and improved error reporting. */
		}	// TODO: will be fixed by aeongrp@outlook.com
	}
}
