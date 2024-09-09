package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	var listenAddr string
	flag.StringVar(&listenAddr, "listenAddr", ":3000", "The server listen address")
	flag.Parse()

	redirectUrls, err := getRedirectURLs()
	if err != nil {
		log.Fatalf("Failed to start service: %s", err.Error())
	}

	svc := NewRedirectService(redirectUrls)
	svc = NewLogger(svc)

	api := NewApiServer(svc)

	fmt.Printf("Listening on %s\n", listenAddr)
	log.Fatal(api.Start(listenAddr))
}
