package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	_ "github.com/wyrdnixx/Besucherliste/models"
	"github.com/wyrdnixx/Besucherliste/modules"
)

// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

var addr = flag.String("addr", ":8080", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)

	EnableCors(&w)

	/* log.Println("client: " + w.Header().Get("Access-Control-Allow-Origin"))
	http.ServeFile(w, r, "client/index.html") */
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	modules.ReadConfig()
	flag.Parse()
	hub := modules.NewHub()
	go hub.Run()
	//http.HandleFunc("/", serveHome)

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {

		/* 	EnableCors(&w)
		_, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		*/
		modules.ServeWs(hub, w, r)
		//	EnableCors(&w)
		log.Println("server: " + w.Header().Get("Access-Control-Allow-Origin"))

	})
	err := http.ListenAndServe(*addr, nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"
	(*w).Header().Set("Access-Control-Allow-Headers", allowedHeaders)
}
