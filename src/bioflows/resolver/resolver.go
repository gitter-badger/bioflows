package resolver

import "strings"

const (
	BIOFLOWS_NAME = "bioflows"
	BIOFLOWS_PIPELINES = "pipelines"
	BIOFLOWS_NODES = "nodes"
	BIOFLOWS_LEADER = "leader"

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

