package executer_services

import (
	"context"
	agentexepb "github.com/headend/agent-executer-service/proto"
	static_config "github.com/headend/share-module/configuration/static-config"
	file_and_directory "github.com/headend/share-module/file-and-directory"
	"github.com/headend/share-module/model"
	"log"
	"time"
)

func (c *agentExeServer) RunUrgentTask(ctx context.Context, in *agentexepb.AgentEXERequest) (*agentexepb.AgentEXEResponse, error) {
	/*
	1. Check agent exits (by pass, gateway do it)
	2. Assign signal number
	3. Send signal
	 */
	// Check agent exists
	// Assign signal number
	// Send signal
	exeResponseData, err2 := SendExeSignal(in, static_config.UrgentTask)
	if err2 != nil {
		return nil, err2
	}
	return &exeResponseData, nil
}

func (c *agentExeServer) RunShell(ctx context.Context, in *agentexepb.AgentEXERequest) (*agentexepb.AgentEXEResponse, error) {
	return nil, nil
}


func SendExeSignal(in *agentexepb.AgentEXERequest, exeType int) (agentexepb.AgentEXEResponse, error) {
	err2 := SendMsgToQueue(in, exeType)
	if err2 != nil {
		log.Fatalln(err2)
		return agentexepb.AgentEXEResponse{}, err2
	}
	// make response data
	var AgentExeResponse []*agentexepb.AgentEXERequest
	AgentExeResponse = append(AgentExeResponse, in)
	exeResponseData := agentexepb.AgentEXEResponse{
		AgentEXEResponseStatus: true,
		Agentexes:           AgentExeResponse,
	}
	return exeResponseData, nil
}

func SendMsgToQueue(in *agentexepb.AgentEXERequest, exeType int) (err error) {
	messageData := model.AgentEXEQueueRequest{
		AgentExeRequest: model.AgentExeRequest{
			AgentIp:        in.AgentIp,
			ExeId:          int(in.ControlId),
		},
		ExeType:         exeType,
		EventTime:       time.Now().Unix(),
	}

	msgSendToQueue, err := messageData.GetJsonString()
	if err != nil {
		return err
	}
	// Do send to message queue
	logFile := file_and_directory.MyFile{Path: static_config.LogPath + "/exemsg.log"}
	logFile.WriteString(msgSendToQueue)
	return nil
}

