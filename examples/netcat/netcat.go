package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	c, err := net.Dial("tcp", "localhost:8080")
	conn, _ := c.(*net.TCPConn)

	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})

	go func() {
		io.Copy(os.Stdout, conn)
		fmt.Println("done")
		done <- struct{}{}
	}()

	mustCopy(conn, os.Stdin)
	conn.CloseWrite()
	<-done // wait for background goroutine to finish
}

func mustCopy(w io.Writer, src io.Reader) {
	if _, err := io.Copy(w, src); err != nil {
		log.Fatal(err)
	}
}
