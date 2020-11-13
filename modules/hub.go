// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package modules

import (
	"encoding/json"
	"fmt"

	_ "github.com/wyrdnixx/Besucherliste/models"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// bradcast messages to all the clients.
	broadcast chan []byte

	// incomming Message
	incomming chan transmitter

	// Register requests from the clients.
	register chan *Client

	// cast message to single client
	singlecast chan transmitter

	// Unregister requests from clients.
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		incomming:  make(chan transmitter),
		singlecast: make(chan transmitter),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			fmt.Printf("New Client registered: %+v\n", client.conn.RemoteAddr())
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				fmt.Printf("client disconnected: %s", client.conn.RemoteAddr())
				delete(h.clients, client)
				close(client.send)
			}
		/* case message := <-h.broadcast:
		fmt.Printf("Broadcasting message: %s ", message)
		for client := range h.clients {
			select {
			case client.send <- message:
			default:
				close(client.send)
				delete(h.clients, client)
			}
		} */

		case message := <-h.broadcast:
			fmt.Printf("Broadcasting message: %s ", message)
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		case tr := <-h.singlecast:
			fmt.Printf("sending single targat message: %s ", tr.Message)
			tr.Client.send <- []byte(tr.Message)
			//client.send <- []byte("result to singel client")

		case incomming := <-h.incomming:
			fmt.Printf("Incomming from client: %s <> %s\n", incomming.Client.conn.RemoteAddr(), incomming.Message)
			//incomming.Client.send <- []byte("Holla")
			b, err := json.Marshal(ConsumeMessage(incomming))
			if err != nil {
				// nothing - b contains statusmessage - error or result
			}
			incomming.Client.send <- []byte(b)

		}

	}
}
