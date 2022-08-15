package mongodb

import (
	"encoding/json"
	"net/http"

	"github.com/jongwon254/Chat-Room/model"
)

// MongoDB Controller

// insert message into database
func InsertMessage(message model.Message) {
	insertMessage(message)
}

// get chat history, set request header and encode into JSON
func GetAllMessages(w http.ResponseWriter, r *http.Request) {
	allMessages := getAll()

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(allMessages)
}

// delete chat history, set request header and encode into JSON
func DeleteAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	count := deleteAll()
	json.NewEncoder(w).Encode(count)
}
