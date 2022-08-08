package router

import (
	"github.com/gorilla/mux"
	"github.com/jongwon254/Chat-Room/mongodb"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/messages", mongodb.GetAllMessages).Methods("GET")
	router.HandleFunc("/api/delete", mongodb.DeleteAll).Methods("DELETE")

	return router
}
