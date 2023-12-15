package main

import (
	"crypto/tls"
	"net/http"
	"os"
	"log"
)

var dl *log.Logger = log.New(os.Stdout, "[Debug]: ", log.Lshortfile)

func main() {
	// load the certificate
	cert, err__load_cert := tls.LoadX509KeyPair("./certs/ssl_cert.pem", "./certs/cert-pkey.pem")
	if err__load_cert != nil {
		dl.Printf("failed to load SSL certificates. %s\n", err__load_cert)
		os.Exit(1)
	}

	// create a TLS config for the HTTPs certificate
	tls_config := &tls.Config{
		Certificates: []tls.Certificate{cert},
		// NOTE: when creating a proxy you can get the `server-name` from here
		// TODO: when creating the proxy, match the `ClientHello` information
		GetConfigForClient: func(c_config *tls.ClientHelloInfo) (*tls.Config, error) {
			dl.Printf("Clint Config\n%v\n", *c_config)
			return nil, nil
		},
	}

	https_server := &http.Server{
		Addr: ":8443",
		TLSConfig: tls_config,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		dl.Printf("got a client request on /\n")
	})

	dl.Printf("listening for https connections on port 8443\n")

	err__start_server := https_server.ListenAndServeTLS("", "")
	if err__start_server != nil {
		dl.Printf("failed to start HTTPS server. %s\n", err__start_server)
		os.Exit(1)
	}
}
