package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/flibustenet/chat/handlers"
)

//go:embed static/*.js
var StaticFS embed.FS

type MiddleWare struct {
	mux http.Handler
}

func (m *MiddleWare) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	m.mux.ServeHTTP(w, r)
	log.Println(r.Method, r.URL.Path, time.Since(start))
}

func midFunc(mux http.Handler) http.Handler {
	return mux
}
func main() {
	mux := http.NewServeMux()
	mux.Handle("/static/", http.FileServer(http.FS(StaticFS)))
	mux.HandleFunc("/events", handlers.Events)
	mux.HandleFunc("/message", handlers.Message)
	mux.HandleFunc("/", handlers.Index)

	mid := &MiddleWare{mux}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Listen :" + port)
	err := http.ListenAndServe(":"+port, mid)
	if err != nil {
		panic(err)
	}
}
