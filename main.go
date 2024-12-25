package main

import (
	"fmt"
	"log"
	"net"
)

const port = ":8888"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			continue
		}

		buf := make([]byte, 1024)

		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read conn error")
			return
		}
		fmt.Println("log: " + string(buf[:n]))
	}
}
