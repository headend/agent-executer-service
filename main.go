package main

import (
	"log"
	"github.com/headend/agent-control-service/control-services"
)


func main()  {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Agent control service")
	control_services.StartServer()
}
