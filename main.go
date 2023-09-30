package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/koshkaj/expensebot/server"
)

func main() {
	s := server.InitServer()
	go func() {
		s.Logger.Fatal(s.Start(s.Config.Server.Port))
	}()
	terminate := make(chan os.Signal, 1)
	signal.Notify(terminate, os.Interrupt, os.Kill)
	<-terminate
	fmt.Println("Shutting down server...")
}
