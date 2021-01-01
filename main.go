package main

import (
	"flag"
	"log"
	"net/http"

	_ "github.com/wyrdnixx/Besucherliste/models"
	"github.com/wyrdnixx/Besucherliste/modules"
)

// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

var addr = flag.String("addr", ":8080", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)

	/*
		if r.URL.Path != "/" {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}
		if r.Method != "GET" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		//http.ServeFile(w, r, "websockets.html")

	*/
	w.Header().Set("contend-type", "application/javasript")
	http.ServeFile(w, r, "client/index.html")
}

func main() {
	modules.ReadConfig()
	flag.Parse()
	hub := modules.NewHub()
	go hub.Run()
	http.HandleFunc("/", serveHome)

	//fs := http.FileServer(http.Dir("client"))
	//http.Handle("/", fs)

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {

		modules.ServeWs(hub, w, r)
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
