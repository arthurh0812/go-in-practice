package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	ln, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("listening on TCP: fail: %v", err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalf("establishing connection: fail: %v", err)
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	input := bufio.NewScanner(conn)
	for input.Scan() {
		echo(conn, input.Text(), 2*time.Second)
	}
	conn.Close()
}

func echo(w io.Writer, shout string, delay time.Duration) {
	fmt.Fprintln(w, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(w, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(w, "\t", strings.ToLower(shout))
}
