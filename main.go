package main

import (
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	server, err := socketio.NewServer(nil)

	if err != nil {
		log.Fatal(err)
	}

	// sockets
	server.On("connection", func(so socketio.Socket) {
		log.Println("New User Connected.")

		// other events
	})

	// http
	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./client")))
	log.Println("Server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
