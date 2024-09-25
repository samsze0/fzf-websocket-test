package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/gorilla/websocket"
)

var apiKey = "test"
var port = 12000

func startClient() {
	dialer := websocket.DefaultDialer

	req, err := http.NewRequest("GET", fmt.Sprintf("ws://localhost:%d", port), nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Fzf-Api-Key", apiKey)

	c, _, err := dialer.DialContext(context.Background(), req.URL.String(), req.Header)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read error:", err)
				return
			}
			log.Printf("received: %s", message)
		}
	}()

	go func() {
		defer wg.Done()
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			message := scanner.Text()
			err := c.WriteMessage(websocket.TextMessage, []byte(message))
			if err != nil {
				log.Println("write error:", err)
				return
			}
		}
		if err := scanner.Err(); err != nil {
			log.Println("input error:", err)
		}
	}()

	wg.Wait()
}