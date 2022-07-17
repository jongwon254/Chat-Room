package main

import (
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	server := socketio.NewServer(nil)

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		return nil
	})

	// http
	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./client")))
	log.Println("Server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
