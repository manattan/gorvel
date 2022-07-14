package main

import (
	"log"

	"github.com/manattan/gorvel/server"
)

func main() {
	s, err := server.NewServer()
	if err != nil {
		log.Fatalln(err)
	}

	err = s.Start(":1323")
	if err != nil {
		log.Fatalln(err)
	}
}