package pipelines

import (
	"fmt"
	viz "github.com/awalterschulze/gographviz"
	"github.com/goombaio/dag"
	"strings"
)



func PreparePipeline(b *BioPipeline) (*BioPipeline,error) {
	//TODO: this function should perform the following tasks
	// 1. Download the tool from the remote repository, in this order (URL , Bioflows Hub)
	// 2. Update the downloaded tool parameters by the newly written parameters.
	return b , nil
}

func CreateGraph(b *BioPipeline) (*dag.DAG,error){
	g := dag.NewDAG()
	processedSteps := make(map[string]*dag.Vertex)
	for _ , step := range b.Steps {
		step.Prepare()
		if len(step.Depends) <= 0{
			vertex := dag.NewVertex(step.ID,step)
			g.AddVertex(vertex)
			processedSteps[step.ID] = vertex
		}else{
			from := step.Depends
			currentVertex := dag.NewVertex(step.ID,step)
			if parentVertex, ok := processedSteps[from]; !ok {
				panic(fmt.Errorf("Unknown Bioflows Step mentioned in %s",step.Name))
			}else{
				g.AddVertex(currentVertex)
				g.AddEdge(parentVertex,currentVertex)
				processedSteps[step.ID] = currentVertex
			}
		}
	}
	return g, nil
}

func ToDotGraph(b *BioPipeline , d *dag.DAG)  (*viz.Graph,error) {
	parents := d.SourceVertices()
	g := viz.NewGraph()

	g.SetName(strings.ReplaceAll(b.Name," ","_"))
	for _ , parent := range parents {
		current := parent.Value.(BioPipeline)
		g.AddNode(b.Name,current.Name,nil)
		if parent.Children.Size() > 0 {
			for _, child := range parent.Children.Values() {
				currentChild := (child.(*dag.Vertex)).Value.(BioPipeline)
				g.AddNode(b.Name,currentChild.Name,nil)
				g.AddEdge(current.Name,currentChild.Name,true, nil)
			}
		}
	}
	return g, nil
}

