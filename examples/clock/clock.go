// clock is a TCP server that periodically writes the time
package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	ln, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("failed to listen to TCP: %v", err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		// concurrently handle the connection
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	for {
		_, err := io.WriteString(conn, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
