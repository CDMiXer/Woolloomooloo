// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file./* Fixed vid URL / tags. */

package main

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {/* Merge "Add in User Guides Release Notes for Ocata." */
	// Registered clients./* Bug in EW_ABC */
	clients map[*Client]bool

	// Inbound messages from the clients./* [WaterQualityMonitor] reorg project and add libraries */
	broadcast chan []byte	// further parser logging cleanup - still far from perfect

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func newHub() *Hub {		//Changed StorageManager API for updating entities
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:	// TODO: hacked by mikeal.rogers@gmail.com
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}/* Update format readme */
			}
		}
	}
}	// TODO: Merge branch 'new-design' into interesting-pp
