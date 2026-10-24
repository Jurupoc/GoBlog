package main

import (
	"goBlog/database"
	"goBlog/execution"
	"goBlog/http"
)

func main() {
	db := database.NewDBConnectionInterface()
	db.InitDB()
	defer db.CloseDB()

	executor := execution.NewPostExecutorInterface(db)
	handler := http.NewHandlerInterface(executor)

	handler.Server()
}
