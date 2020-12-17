package main

import (
	"log"
	"github.com/headend/agent-executer-service/execute-services"
)


func main()  {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Agent control service")
	execute_services.StartServer()
}
