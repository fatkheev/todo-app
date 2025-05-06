package main

import (
	"log"
	"todo-app"
	"todo-app/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRouts()); err != nil {
		log.Fatalf("error start server: %s", err.Error())
	}
}