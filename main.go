package main

import (
	"fmt"
	"github.com/nats-io/stan.go"
	"io"
	"os"
)

func main() {

	clusterID := "cluster"
	clientID := "orders"

	// !!!read about stan adn conn
	sc, _ := stan.Connect(clusterID, clientID)

	// reading file
	// !!! read about reading file
	f, err := os.Open(os.Args[1]) //?
	if err != nil {
		fmt.Println(err)
	}

	data,err := io.ReadAll(f)
	if err != nil {
		fmt.Println(err)
	}

	// Simple Synchronous Publisher
	sc.Publish("orders", []byte(data)) // does not return until an ack has been received from NATS Streaming

	// Simple Async Subscriber
	// do i need it ??
	/*
	sub, _ := sc.Subscribe("foo", func(m *stan.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})

	// Unsubscribe
	sub.Unsubscribe()*/

	// Close connection
	sc.Close()

}