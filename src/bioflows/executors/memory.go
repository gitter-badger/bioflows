package executors

import (
	"github.com/goombaio/dag"
)

type ExecutorMemory struct {
	WaitQueue chan *dag.Vertex
	elements map[string] *dag.Vertex

}
func (e *ExecutorMemory) SetUp(){
	e.WaitQueue = make(chan *dag.Vertex,5)
	e.elements = make(map[string]*dag.Vertex)
}
func (e *ExecutorMemory) Size() int {
	return len(e.WaitQueue)
}
func (e *ExecutorMemory) AddToMemory(v *dag.Vertex) bool{
	if _ , ok := e.elements[v.ID]; ok {
		return false
	}
	e.elements[v.ID] = v
	e.WaitQueue <- v
	return true
}
func (e *ExecutorMemory) PopFromMemory() (*dag.Vertex,bool) {
	item , ok := <- e.WaitQueue
	if ok {
		delete(e.elements,item.ID)
	}
	return item , ok
}
