package main

import (
	"log"
	"github.com/headend/agent-executer-service/executer-services"
)


func main()  {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Agent control service")
	executer_services.StartServer()
}
