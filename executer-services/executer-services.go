package executer_services

import (
	"fmt"
	"github.com/headend/share-module/configuration"
	agentexepb "github.com/headend/agent-executer-service/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type agentExeServer struct {
	config *configuration.Conf
}


func StartServer()  {
	var config configuration.Conf

	config.LoadConf()
	listenAddr := fmt.Sprintf("%s:%d", config.RPC.Agentrunner.Host, config.RPC.Agentrunner.Port)
	ln, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	rpcServer := grpc.NewServer()
	agentexepb.RegisterAgentEXEServiceServer(rpcServer, &agentExeServer{config:&config})
	if rpcServer == nil {
		log.Fatalf("failed to register server: %v", err)
	}
	if err := rpcServer.Serve(ln); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}