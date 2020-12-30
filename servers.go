package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"sync"
	"time"
)

func slowServer(wg *sync.WaitGroup) {
	s := &http.Server{
		Addr: ":8080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(2 * time.Second)
			w.Write([]byte("Slow response"))
		}),
		// this is a trick to let us know that the server has finished launching.
		// if we don't do this, the servers might start after the clients attempt
		// to connect to them, causing an error.
		BaseContext: func(net.Listener) context.Context {
			defer wg.Done()
			return context.Background()
		},
	}
	log.Fatal(s.ListenAndServe())
}

func errServer(wg *sync.WaitGroup) {
	s := &http.Server{
		Addr: ":9090",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Query().Get("error") == "true" {
				w.Write([]byte("error"))
				return
			}
			w.Write([]byte("ok"))
		}),
		// this is a trick to let us know that the server has finished launching.
		// if we don't do this, the servers might start after the clients attempt
		// to connect to them, causing an error.
		BaseContext: func(net.Listener) context.Context {
			defer wg.Done()
			return context.Background()
		},
	}
	log.Fatal(s.ListenAndServe())
}
