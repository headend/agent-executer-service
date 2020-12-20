package executer_services

import (
	"context"
	agentexepb "github.com/headend/agent-executer-service/proto"
	queueServer "github.com/headend/share-module/MQ"
	static_config "github.com/headend/share-module/configuration/static-config"
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
	exeResponseData, err2 := SendExeSignal(c, in, static_config.UrgentTask)
	if err2 != nil {
		return nil, err2
	}
	return &exeResponseData, nil
}

func (c *agentExeServer) RunShell(ctx context.Context, in *agentexepb.AgentEXERequest) (*agentexepb.AgentEXEResponse, error) {
	return nil, nil
}


func SendExeSignal(c *agentExeServer, in *agentexepb.AgentEXERequest, exeType int) (agentexepb.AgentEXEResponse, error) {
	err2 := SendMsgToQueue(c, in, exeType)
	if err2 != nil {
		log.Fatalln(err2)
		return agentexepb.AgentEXEResponse{}, err2
	}
	// make response data
	var AgentExeResponse []*agentexepb.AgentEXERequest
	AgentExeResponse = append(AgentExeResponse, in)
	exeResponseData := agentexepb.AgentEXEResponse{
		Status: agentexepb.AgentEXEResponseStatus_SUCCESS,
		Agentexes:           AgentExeResponse,
	}
	return exeResponseData, nil
}

func SendMsgToQueue(c *agentExeServer, in *agentexepb.AgentEXERequest, exeType int) (err error) {
		messageData := model.AgentEXEQueueRequest{
		AgentExeSingleRequest: model.AgentExeSingleRequest{
			AgentId:    in.AgentId,
			ExeType:    exeType,
			ExeId:      0,
			TunnelData: nil,
		},
		ExeType:               exeType,
		EventTime:             time.Now().Unix(),
	}

	msgSendToQueue, err := messageData.GetJsonString()
	if err != nil {
		return err
	}
	// Do send to message queue

	var msgQueueServer queueServer.MQ
	//defer msgQueueServer.CloseProducer()
	msgQueueServer.PushMsgByTopic(c.config, msgSendToQueue,c.config.MQ.CommandTopic)
	if msgQueueServer.Err != nil {
		log.Println(msgQueueServer.Err.Error())
		return msgQueueServer.Err
	}
	return nil
}

