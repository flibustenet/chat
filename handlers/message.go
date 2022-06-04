package handlers

import (
	"fmt"
	"net/http"

	"github.com/flibustenet/chat/broker"
)

func Message(w http.ResponseWriter, r *http.Request) {
	msg := r.FormValue("message")
	broker.Broker.Send(msg)
	fmt.Fprintf(w, "<input name='message' id='message'/>")
}
