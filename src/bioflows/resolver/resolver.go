package resolver

import "strings"

const (
	BIOFLOWS_NAME = "BIOFLOWS"
	BIOFLOWS_PIPELINES = "PIPELINES"
	BIOFLOWS_NODES = "NODES"
	BIOFLOWS_LEADER = "LEADER"

)

func ResolveNodeKey(nodeId string) string {
	return strings.Join([]string{BIOFLOWS_NAME,BIOFLOWS_NODES,nodeId},"/")
}
func ResolveToolKey(toolId string , pipelineId string) string {
	return strings.Join([]string{BIOFLOWS_NAME,BIOFLOWS_PIPELINES,pipelineId,toolId},"/")
}

func ResolveLeaderKey() string {
	return strings.Join([]string{BIOFLOWS_NAME,BIOFLOWS_NODES,BIOFLOWS_LEADER},"/")
}

func ResolvePipelineKey(pipelineId string) string {
	return strings.Join([]string{BIOFLOWS_NAME,BIOFLOWS_PIPELINES,pipelineId},"/")
}

