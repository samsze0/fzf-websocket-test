package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/websocket"
)

var apiKey = "test"
var port = 12010

var clients = make(map[*websocket.Conn]string)
var broadcast = make(chan string)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, http.Header{
		"Fzf-Api-Key": []string{apiKey},
	})
	if err != nil {
		log.Fatal(err)
	}

	defer ws.Close()

	clients[ws] = fmt.Sprintf("Client%d", rand.Intn(100))

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			delete(clients, ws)
			break
		}

		log.Printf("Received message from %s: %s", clients[ws], string(msg))
	}
}

func handleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func broadcastMessage(message string) {
	broadcast <- message
}

func startServer(port int) {
	var listener net.Listener
	var err error

	listener, err = net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/", handleConnections)

	go handleMessages()

	log.Println("websocket server started on port", port)

	err = http.Serve(listener, mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func main() {
	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			log.Println("waiting for input...")

			line, err := reader.ReadString('\n')
			if err != nil {
				log.Printf("error while reading line from stdin: %v", err)
				break
			}

			line = strings.TrimSpace(line) // Remove trailing newline

			log.Println("broadcasting message: ", line)
			broadcastMessage(line)
		}
	}()

	startServer(port)
}
