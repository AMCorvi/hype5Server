package main

import (
	"fmt"
	"net/http"
)

type home struct {
	welcomeMessage string
}

func (h home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("%s", h.welcomeMessage)))
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hype is on the Server"))
	})

	http.Handle("/cardi", &home{welcomeMessage: "Making Shmoney moves"})

	http.ListenAndServe(":8000", nil)
}
