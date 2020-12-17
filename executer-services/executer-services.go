package executer_services

import (
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
	ln, err := net.Listen("tcp", "0.0.0.0:5000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	rpcServer := grpc.NewServer()
	agentexepb.RegisterAgentEXEServiceServer(rpcServer, &agentExeServer(config:&config))
	if rpcServer == nil {
		log.Fatalf("failed to register server: %v", err)
	}
	if err := rpcServer.Serve(ln); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}


}