package router

import (
	"github.com/gorilla/mux"
	"github.com/jongwon254/Chat-Room/mongodb"
)

// API endpoints for chat history
func Router() *mux.Router {
	router := mux.NewRouter()

	// GET
	router.HandleFunc("/api/messages", mongodb.GetAllMessages).Methods("GET")

	// DELETE
	router.HandleFunc("/api/delete", mongodb.DeleteAll).Methods("DELETE", "OPTIONS")

	return router
}
