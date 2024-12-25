package main

import (
	"log"
	"net"
)

const (
	transport = "tcp"
	port      = ":8888"
)

func main() {
	var err error
	conn, err := net.Dial(transport, port)
	if err != nil {
		log.Fatal(err)
	}
	message := "helloooo"
	_, err = conn.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}

}
