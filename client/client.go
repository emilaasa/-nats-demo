package main

import (
	"fmt"
	"github.com/nats-io/stan.go"
)

func main() {
	fmt.Println("ITSA MEE")
	sc, err := stan.Connect("test-cluster", "hedgebog")
	if err != nil {
		panic(err)
	}
	defer sc.Close()

	//sub, err := sc.Subscribe("foo", func(m *stan.Msg) {
	//	fmt.Printf("RECV: %s\n", m.Data)
	//}, stan.DeliverAllAvailable())

	sub, _ := sc.QueueSubscribe("foo", 
		)
	sub.Close()
}