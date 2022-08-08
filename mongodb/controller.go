package mongodb

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jongwon254/Chat-Room/model"
)

// db controller

// insert in backend
func InsertMessage(message model.Message) {
	insertMessage(message)
}

// get via endpoint
func GetAllMessages(w http.ResponseWriter, r *http.Request) {
	allMessages := getAll()
	fmt.Println("messages", allMessages)

	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	json.NewEncoder(w).Encode(allMessages)
}

// delete via endpoint
func DeleteAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAll()
	json.NewEncoder(w).Encode(count)
}
