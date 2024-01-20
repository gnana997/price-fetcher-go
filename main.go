package main

import (
	"flag"
)

func main() {

	listenAddr := flag.String("listenAddr", ":3000", "listen address")
	flag.Parse()

	svc := NewLogginServe(&priceFetcher{})

	server := NewJSONAPIServer(*listenAddr, svc)

	server.Run()
}
