package main

import (
	"net/http"
	"time"
)

func slowServer() {
	s := &http.Server{
		Addr: ":8080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(2 * time.Second)
			w.Write([]byte("Slow response"))
		}),
	}
	s.ListenAndServe()
}

func errServer() {
	s := &http.Server{
		Addr: ":9090",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Query().Get("error") == "true" {
				w.Write([]byte("error"))
				return
			}
			w.Write([]byte("ok"))
		}),
	}
	s.ListenAndServe()
}
