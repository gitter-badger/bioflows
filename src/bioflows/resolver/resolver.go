package resolver

import "strings"

const (
	BIOFLOWS_NAME = "bioflows"
	BIOFLOWS_META = "meta"
	BIOFLOWS_PIPELINES = "pipelines"
	BIOFLOWS_NODES = "nodes"
	BIOFLOWS_LEADER = "leader"

)

func ResolveNodeKey(nodeId string) string {
	//Node Key: bioflows/nodes/%s
	return strings.Join([]string{BIOFLOWS_NAME,BIOFLOWS_NODES,nodeId},"/")
}
func ResolveMetaDataForLeader() string {
	return strings.Join([]string{BIOFLOWS_NAME,BIOFLOWS_META,BIOFLOWS_LEADER},"/")
}
func ResolveMetaDataForNode(nodeId string) string {
	return strings.Join([]string{BIOFLOWS_NAME,BIOFLOWS_META,nodeId},"/")
}
func ResolveToolKey(toolId string , pipelineId string) string {
	// Tool Key: bioflows/pipelines/%pId/%tId
	return strings.Join([]string{BIOFLOWS_NAME,BIOFLOWS_PIPELINES,pipelineId,toolId},"/")
}

func ResolveLeaderKey() string {
	//Leader Key: bioflows/nodes/leader
	return strings.Join([]string{BIOFLOWS_NAME,BIOFLOWS_NODES,BIOFLOWS_LEADER},"/")
}

func ResolvePipelineKey(pipelineId string) string {
	return strings.Join([]string{BIOFLOWS_NAME,BIOFLOWS_PIPELINES,pipelineId},"/")
}

