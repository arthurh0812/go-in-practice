package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	_ "crypto/sha256"
	"io"
	"log"
	"net"
	"strconv"
)

var privKey *rsa.PrivateKey

func main() {
	privKey, _ = rsa.GenerateKey(rand.Reader, 2048)

	addr := "127.0.0.1:8080"
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen to TCP: %v", err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalf("failed to establish connection: %v", err)
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	data := make([]byte, 1024)
	n, err := conn.Read(data)
	if err != nil && err != io.EOF {
		log.Fatalf("failed to read input: %v", err)
	}
	first := string(data[:n])
	log.Printf("Read %d bytes.", n)
	log.Print(first)

	pubKey := privKey.PublicKey

	if first == "public" {
		e, n := strconv.FormatInt(int64(pubKey.E), 10), pubKey.N.Text(10)
		conn.Write([]byte(e + " " + n + "\n"))
	}

	n, err = conn.Read(data)
	if err != nil {
		log.Fatalf("failed to read input: %v", err)
	}
	cipher := data[:n]

	plain, err := privKey.Decrypt(nil, cipher, &rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		log.Fatalf("failed to decrypt cipher text: %v", err)
	}

	log.Printf("plain text: %s", plain)
}
