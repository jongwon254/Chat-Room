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

	// broadcast welcome message
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
			date := time.Now().String()
			user := "User " + s.ID()
			message := model.Message{ID: primitive.NewObjectID(), User: user, Text: msg, Date: date}
			mongodb.InsertMessage(message)

			// broadcast
			msg := user + ": " + msg
			server.BroadcastToRoom("", "chat-room", "chat message", msg)
		}

	})

	server.OnError("/", func(s socketio.Conn, e error) {
		log.Println("Error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("Closed.", reason)
	})

	go server.Serve()
	defer server.Close()

	// http api
	go func() {
		fmt.Println("Server starting...")
		r := router.Router()
		log.Fatal(http.ListenAndServe(":8080", r))
		fmt.Println("Listening on port 8080...")
	}()

	// http websocket
	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./client")))
	log.Println("Server on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))

}
