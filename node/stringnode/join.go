package stringnode

import "strings"

// Join joins data and makes string slice of length 1.
type Join struct {
	done bool

	data  []string
	error error

	inputs []Node // Join should have 1 input.
	parm   JoinParm
}

type JoinParm struct {
	with string
}

// NewJoin creates a new Join node.
func NewJoin(parm JoinParm) *Join {
	n := &Join{
		inputs: make([]Node, 1),
		parm:   parm,
	}
	return n
}

// Type is a type name of the node.
func (n *Join) Type() string {
	return "Join"
}

// Inputs returns it's inputs.
func (n *Join) Inputs() []Node {
	return n.inputs
}

// AddInput set or replaces the first input Node.
func (n *Join) AddInput(in Node) {
	n.inputs[0] = in
}

// Result returns it's calculated data.
func (n *Join) Result() ([]string, error) {
	if !n.done {
		n.join()
	}
	n.done = true

	return n.data, n.error
}

// join joins it's data.
func (n *Join) join() {
	if n.inputs[0] == nil {
		n.error = NewError(n, "should have 1 input")
	}

	inData, err := n.inputs[0].Result()
	if err != nil {
		n.error = err
		return
	}
	if inData == nil {
		n.error = NewError(n, "first input's data should not nil")
		return
	}

	if len(inData) == 0 {
		n.data = []string{}
		return
	}

	n.data = []string{strings.Join(inData, n.parm.with)}
}
