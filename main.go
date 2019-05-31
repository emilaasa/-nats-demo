package main

import (
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
)

func main() {
	fmt.Println("moo")

	sc, err := stan.Connect("test-cluster", "hedgehog")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer sc.Close()

	err = sc.Publish("foo", []byte("Hello World 1"))
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	sub, _ := sc.Subscribe("foo", func(m *stan.Msg) {
		fmt.Printf("RECV: %s\n", m.Data)
	})

	_ = sub.Unsubscribe()
	go func() {

	}()
}
