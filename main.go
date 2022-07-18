package main

import (
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	server := socketio.NewServer(nil)

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		log.Println("User", s.ID(), "connected.")
		s.Join("chat-room")
		return nil
	})

	server.OnEvent("/", "welcome", func(s socketio.Conn, msg string) {
		server.BroadcastToRoom("", "chat-room", "welcome", msg)
	})

	server.OnEvent("/", "msg", func(s socketio.Conn, msg string) {
		//s.Emit("chat message", msg)
		if msg == "disconnect" {
			server.BroadcastToRoom("", "chat-room", "chat message", "User Disconnected.")
			s.Close()
			server.Close()
		} else {
			msg := "User " + s.ID() + ": " + msg
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

	// http
	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./client")))
	log.Println("Server on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))

}
