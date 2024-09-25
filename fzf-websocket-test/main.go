package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <client|server>")
		os.Exit(1)
	}

	mode := os.Args[1]

	switch mode {
	case "client":
		startClient()
	case "server":
		// startServer()
	default:
		fmt.Println("Invalid argument. Use 'client' or 'server'.")
		os.Exit(1)
	}
}
