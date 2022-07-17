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
		return nil
	})

	server.OnEvent("/", "msg", func(s socketio.Conn, msg string) {
		log.Println(msg)
	})

	go server.Serve()
	defer server.Close()

	// http
	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./client")))
	log.Println("Server on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))

}
