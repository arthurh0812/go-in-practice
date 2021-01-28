package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"log"
	"math/big"
	"net"
	"strconv"
	"strings"
)

func main() {
	addr := "127.0.0.1:8080"

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatalf("failed to establish connection: %v", err)
	}
	defer conn.Close()

	conn.Write([]byte("public"))

	data := make([]byte, 1024)

	m, err := conn.Read(data)
	if err != nil {
		log.Fatalf("failed to read data: %v", err)
	}

	input := string(data[:m])

	buf := bytes.NewBufferString(input)

	eString, err := buf.ReadString(' ')
	if err != nil {
		log.Fatalf("failed to read until next space: %v", err)
	}
	eString = strings.TrimSuffix(eString, " ")

	nString, err := buf.ReadString('\n')
	if err != nil {
		log.Fatalf("failed to read until next space: %v", err)
	}
	nString = strings.TrimSuffix(nString, "\n")

	var e int
	if v, err := strconv.ParseFloat(eString, 10); err == nil {
		e = int(v)
	}
	var n = new(big.Int)
	n, ok := n.SetString(nString, 10)
	if !ok {
		log.Fatalf("failed to convert to big int")
	}

	pubKey := &rsa.PublicKey{
		E: e,
		N: n,
	}

	message := "super secret message"

	cipher, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		pubKey,
		[]byte(message),
		nil,
	)
	if err != nil {
		log.Fatalf("failed to encrypt the message: %v", err)
	}

	conn.Write(cipher)
}
