package stringnode

import (
	"fmt"
	"strings"
)

type ReplaceNode struct {
	done   bool
	result string
	error  error

	inputs []Node

	parm ReplaceNodeParm
}

type ReplaceNodeParm struct {
	from string
	to   string
	n    int
}

func NewReplaceNode(parm ReplaceNodeParm) *ReplaceNode {
	return &ReplaceNode{
		inputs: make([]Node, 0),
		parm:   parm,
	}
}

func (n *ReplaceNode) Inputs() []Node {
	return n.inputs
}

func (n *ReplaceNode) AddInput(in Node) {
	n.inputs = append(n.inputs, in)
}

func (n *ReplaceNode) Result() (string, error) {
	if !n.done {
		n.replace()
	}
	n.done = true

	return n.result, n.error
}

func (n *ReplaceNode) replace() {
	if n.inputs == nil {
		n.error = fmt.Errorf("input is nil")
		return
	}
	if len(n.inputs) != 1 {
		n.error = fmt.Errorf("should have 1 input")
		return
	}

	in, err := n.inputs[0].Result()
	if err != nil {
		n.error = err
		return
	}

	n.result = strings.Replace(in, n.parm.from, n.parm.to, n.parm.n)
	return
}
