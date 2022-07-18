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
		log.Println("User" + s.ID() + "connected.")
		s.Join("chat-room")
		return nil
	})

	server.OnEvent("/", "welcome", func(s socketio.Conn, msg string) {
		server.BroadcastToRoom("", "chat-room", "welcome", msg)
	})

	server.OnEvent("/", "msg", func(s socketio.Conn, msg string) {
		//s.Emit("chat message", msg)
		server.BroadcastToRoom("", "chat-room", "chat message", msg)
	})

	go server.Serve()
	defer server.Close()

	// http
	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./client")))
	log.Println("Server on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))

}
