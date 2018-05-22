package main

import (
	"log"
	"net/http"
	"github.com/gorilla/websocket"
	)

var clients = make(map[*websocket.Conn] bool)
var broadcast = make(chan Message)
var upgrader = websocket.Upgrader{}

type Message struct {
        Email    string `json:"email"`
        Username string `json:"username"`
        Message  string `json:"message"`
}

var messages = make([]Message, 0, 3);

func hello_func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello\n"))
}

func ws_connect(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer ws.Close()

	clients[ws] = true

	for client := range clients {
		if client == ws {
			for _, msg := range messages {
				client.WriteJSON(msg)
			}
			break
		}
	}
	
	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if msg.Message != "" {
			messages = append(messages, msg);
		}
		if err != nil {
			log.Println("error: %v", err)
			delete(clients, ws)
			break
		}
		broadcast <- msg
	}
}

func handle_message() {
	log.Printf("Run handle message")
    for {
    	log.Println("Waiting for clients")
        msg := <-broadcast
        for client := range clients {
        	log.Println("Sending to clients")
            err := client.WriteJSON(msg)
            if err != nil {
                log.Printf("error: %v", err)
                client.Close()
                delete(clients, client)
            }
        }
    }
}

func main(){
	log.Println("Server start on :8080 port")

	http.HandleFunc("/", hello_func)
	http.HandleFunc("/ws", ws_connect)

	upgrader.CheckOrigin = func(r *http.Request) bool {
		// allow all connections by default
		return true
	}

	go handle_message()

	err := http.ListenAndServe(":8080", nil)
	if err != nil{
		log.Fatal("ListenAndServe ", err)
	}
}