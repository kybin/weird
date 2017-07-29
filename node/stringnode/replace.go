package stringnode

import (
	"fmt"
	"strings"
)

type Replace struct {
	done   bool
	result []string
	error  error

	inputs []Node

	parm ReplaceParm
}

type ReplaceParm struct {
	from string
	to   string
	n    int
}

func NewReplace(parm ReplaceParm) *Replace {
	return &Replace{
		inputs: make([]Node, 0),
		parm:   parm,
	}
}

func (n *Replace) Inputs() []Node {
	return n.inputs
}

func (n *Replace) AddInput(in Node) {
	n.inputs = append(n.inputs, in)
}

func (n *Replace) Result() ([]string, error) {
	if !n.done {
		n.replace()
	}
	n.done = true

	return n.result, n.error
}

// VeiledResult implements node.Node.
func (n *Replace) VeiledResult() (interface{}, error) {
	data, err := n.Result()
	return interface{}(data), err
}

func (n *Replace) replace() {
	if n.inputs == nil {
		n.error = fmt.Errorf("input is nil")
		return
	}
	if len(n.inputs) != 1 {
		n.error = fmt.Errorf("should have 1 input")
		return
	}

	inData, err := n.inputs[0].Result()
	if err != nil {
		n.error = err
		return
	}

	n.result = make([]string, len(inData))
	for i := range inData {
		n.result[i] = strings.Replace(inData[i], n.parm.from, n.parm.to, n.parm.n)
	}
	return
}
