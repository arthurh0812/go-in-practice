package main

import (
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
)

func main() {
	config := tls.Config{
		Rand:       rand.Reader,
		ClientAuth: tls.RequestClientCert,
		ClientCAs:  &x509.CertPool{},
	}
}
