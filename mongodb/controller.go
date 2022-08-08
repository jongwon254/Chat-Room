package mongodb

import (
	"fmt"

	"github.com/jongwon254/Chat-Room/model"
)

// db controller
func InsertMessage(message model.Message) {
	insertMessage(message)
	fmt.Println("inserted:", message.ID)
}

func GetAllMessages() {
	allMessages := getAll()
	fmt.Println(allMessages)
}
