package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/flibustenet/chat/broker"
)

func Events(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)

	if !ok {
		log.Println("ERROR EVENT Flusher")
		return
	}
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("X-Accel-Buffering", "no")

	recep := make(chan string)
	broker.Broker.Add(recep)
	notify := w.(http.CloseNotifier).CloseNotify()
	go func() {
		<-notify
		broker.Broker.Remove(recep)
	}()
	for msg := range recep {
		fmt.Fprintf(w, "event: evt\n")
		fmt.Fprintf(w, "data: %s<br>\n\n", msg)
		flusher.Flush()
	}
}
