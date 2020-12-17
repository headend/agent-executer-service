package executer_services

import (
	"context"
	agentexepb "github.com/headend/agent-executer-service/proto"
)

func (c *agentExeServer) RunUrgentTask(ctx context.Context, in *agentexepb.AgentEXERequest) (*agentexepb.AgentEXEResponse, error) {
	/*
	1. Check agent exits
	2. Assign signal number
	3. Send signal
	 */
	// Check agent exists
	// Assign signal number
	// Send signal
	return nil, nil
}

func (c *agentExeServer) RunShell(ctx context.Context, in *agentexepb.AgentEXERequest) (*agentexepb.AgentEXEResponse, error) {
	return nil, nil
}
