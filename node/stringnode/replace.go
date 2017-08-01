package stringnode

import "strings"

type Replace struct {
	done bool

	data  []string
	error error

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

// Type is a type name of the node.
func (n *Replace) Type() string {
	return "Replace"
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

	return n.data, n.error
}

func (n *Replace) replace() {
	if n.inputs == nil {
		n.error = NewError(n, "input is nil")
		return
	}
	if len(n.inputs) != 1 {
		n.error = NewError(n, "should have 1 input")
		return
	}

	inData, err := n.inputs[0].Result()
	if err != nil {
		n.error = err
		return
	}

	n.data = make([]string, len(inData))
	for i := range inData {
		n.data[i] = strings.Replace(inData[i], n.parm.from, n.parm.to, n.parm.n)
	}
	return
}
