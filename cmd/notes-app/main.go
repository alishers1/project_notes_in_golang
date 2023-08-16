package main

import (
	"log"
	"notes/internal/handlers"
)

func main() {
	if err := handlers.InitRoutes(); err != nil {
		log.Println(err)
		return
	}
}