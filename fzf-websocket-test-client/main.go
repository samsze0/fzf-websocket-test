package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var apiKey = "test"
var port = 12010

func startClient(port int) {
	dialer := websocket.DefaultDialer

	req, err := http.NewRequest("GET", fmt.Sprintf("ws://localhost:%d", port), nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("API-KEY", apiKey)

	c, _, err := dialer.DialContext(context.Background(), req.URL.String(), req.Header)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("err:", err)
			return
		}
		log.Printf("received: %s", message)
	}
}

func main() {
	startClient(port)
}
