package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	socketio "github.com/googollee/go-socket.io"
	"github.com/jongwon254/Chat-Room/model"
	"github.com/jongwon254/Chat-Room/mongodb"
	"github.com/jongwon254/Chat-Room/router"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {

	// new socketio websocket
	server := socketio.NewServer(nil)

	// new user connected
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		log.Println("User", s.ID(), "connected.")
		s.Join("chat-room")
		return nil
	})

	// broadcast welcoming message
	server.OnEvent("/", "welcome", func(s socketio.Conn, msg string) {
		server.BroadcastToRoom("", "chat-room", "welcome", msg)
	})

	// broadcast chat message or close chat
	server.OnEvent("/", "msg", func(s socketio.Conn, msg string) {
		if msg == "disconnect" {
			server.BroadcastToRoom("", "chat-room", "chat message", "User Disconnected.")
			s.Close()
			server.Close()
		} else {
			// insert into db
			date := time.Now().Format(time.RFC822)
			fmt.Println(date)
			user := "User " + s.ID()
			message := model.Message{ID: primitive.NewObjectID(), User: user, Text: msg, Date: date}
			mongodb.InsertMessage(message)

			// broadcast
			msg := user + ": " + msg
			server.BroadcastToRoom("", "chat-room", "chat message", msg)
		}

	})

	// Exceptions
	server.OnError("/", func(s socketio.Conn, e error) {
		log.Println("Error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("Closed.", reason)
	})

	// start and close server
	go server.Serve()
	defer server.Close()

	// HTTP API service for connecting to MongoDB (port 8080)
	go func() {
		fmt.Println("Server starting on port 8080...")
		r := router.Router()
		log.Fatal(http.ListenAndServe(":8080", r))
	}()

	// HTTP Websocket for Chat Room (port 3000)
	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./client")))
	log.Println("Server on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))

}
