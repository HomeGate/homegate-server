package main

import (
	"log"
	"open-home-graph/server"
)

func main() {

	e := server.CreateEndpoint()

	log.Panicln(e.Start(":5050"))
}
