package main

import (
	"embed"
	"fmt"
	"net/http"

	"github.com/flibustenet/chat/handlers"
)

//go:embed static/*.js
var StaticFS embed.FS

func main() {

	http.Handle("/static/", http.FileServer(http.FS(StaticFS)))
	http.HandleFunc("/events", handlers.Events)
	http.HandleFunc("/message", handlers.Message)
	http.HandleFunc("/", handlers.Index)

	fmt.Println("Listen :8888")
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
