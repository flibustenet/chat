package broker

import (
	"log"
)

var Broker *broker

type broker struct {
	Input chan string
	Subs  map[chan string]bool
}

func (b broker) Add(c chan string) {
	b.Subs[c] = true
	log.Println(b.Subs)
}
func (b broker) Remove(c chan string) {
	log.Printf("delete %v", c)
	delete(b.Subs, c)
	log.Println(b.Subs)
}
func (b broker) Send(msg string) {
	for k, _ := range Broker.Subs {
		log.Printf("envoi %s vers %v", msg, k)
		go func(ch chan string) {
			ch <- msg
		}(k)
	}
}

func init() {
	Broker = &broker{}
	Broker.Input = make(chan string)
	Broker.Subs = map[chan string]bool{}
}
