package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/koshkaj/expensebot/server"
)

func main() {
	go func() {
		log.Fatal(server.InitServer())
	}()
	terminate := make(chan os.Signal, 1)
	signal.Notify(terminate, os.Interrupt, os.Kill)
	<-terminate
	log.Println("Shutting down server...")
}
