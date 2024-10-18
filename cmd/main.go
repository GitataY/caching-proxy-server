package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	port := flag.Int("port", 3000, "port to run the server on")
	origin := flag.String("origin", "", "origin server url where requests are fowarded")
	clearCache := flag.Bool("clear-cache", false, "clear cache")

	flag.Parse()

	if *clearCache {
		fmt.Println("Clearing cache")
		// function to clear cache here TODO
		os.Exit(0)
	}

	if *origin == "" {
		log.Fatal("Origin server url must be provided")
	}

	fmt.Printf("Starting caching proxy on port %d, forwarding to origin: %s\n", *port, *origin)
}
