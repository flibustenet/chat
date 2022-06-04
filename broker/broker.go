package broker

import (
	"fmt"
	"log"
	"time"
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
		k <- msg
	}
}

func init() {
	Broker = &broker{}
	Broker.Input = make(chan string)
	Broker.Subs = map[chan string]bool{}
	go func() {
		for i := 0; ; i++ {
			log.Println("Send x")
			Broker.Input <- fmt.Sprintf("tick %d", i)
			time.Sleep(time.Second)
		}
	}()
	//	go run()
}

func run() {
	for {
		select {
		case m := <-Broker.Input:
			log.Println("recoit ", m)
			for k, _ := range Broker.Subs {
				log.Println("Envoi")
				k <- m
			}
		}
	}
}
