package main

import (
	"embed"
	"fmt"
	"net/http"
	"os"

	"github.com/flibustenet/chat/handlers"
)

//go:embed static/*.js
var StaticFS embed.FS

func main() {

	http.Handle("/static/", http.FileServer(http.FS(StaticFS)))
	http.HandleFunc("/events", handlers.Events)
	http.HandleFunc("/message", handlers.Message)
	http.HandleFunc("/", handlers.Index)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Listen :" + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
