package store

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"
)

// Upgrader settings to get the *Conn type
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true }, // To change later
}

// Client represents a connected user
type Client struct {
	ID   string
	Conn *websocket.Conn
	Send chan []byte
}

// Manage all active clients and broadcasts
type Hub struct {
	Clients    map[*Client]bool
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
	mux        sync.Mutex
}

var hub = Hub{
	Clients:    make(map[*Client]bool),
	Broadcast:  make(chan []byte),
	Register:   make(chan *Client),
	Unregister: make(chan *Client),
}

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	// TO ADD websockets auth using tokens
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		println("Upgrade error: ", err)
		return
	}

	client := &Client{
		ID:   r.RemoteAddr,
		Conn: conn,
		Send: make(chan []byte),
	}
	hub.Register <- client

	go client.read()
	go client.write()
}

func (c *Client) read() {
	defer func() {
		hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			println("Read Error: ", err)
			return
		}
		hub.Broadcast <- message
	}
}

func (c *Client) write() {
	defer c.Conn.Close()
	for message := range c.Send {
		if err := c.Conn.WriteMessage(websocket.TextMessage, message); err != nil {
			println("Error Writing: ", err)
			break
		}
	}
}

func HubRunner() {
	for {
		select {
		case client := <-hub.Register:
			hub.mux.Lock()
			hub.Clients[client] = true
			hub.mux.Unlock()
			println("Client Connected 🤝: %s", client.ID)
		case client := <-hub.Unregister:
			hub.mux.Lock()
			if _, ok := hub.Clients[client]; ok {
				delete(hub.Clients, client)
				close(client.Send)
				println("Client Disonnected 🚫: %s", client.ID)
			}
			hub.mux.Unlock()
		case message := <-hub.Broadcast:
			hub.mux.Lock()
			for client := range hub.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(hub.Clients, client)
				}
			}
			hub.mux.Unlock()
		}
	}
}

func CreateMessage(db *sql.DB, message MessagePayload) error {
	_, err := db.Exec("INSERT INTO messages (sender_id, receiver_id, content, timestamp) VALUES ($1, $2, $3, $4)",
		message.SenderID, message.ReceiverID, message.Content, message.Timestamp)
	return err
}

var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr: "localhost:6379", // redis server addres
})

func SendMessage(sender, receiver, message string) error {
	chatStream := fmt.Sprintf("chat:%s:%s", sender, receiver)
	_, err := rdb.XAdd(ctx, &redis.XAddArgs{
		Stream: chatStream,
		Values: map[string]interface{}{
			"sender":  sender,
			"message": message,
			"time":    time.Now().Unix(),
		},
	}).Result()
	return err
}
