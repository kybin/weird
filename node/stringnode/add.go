package stringnode

// Add appends new strings to it's input data.
type Add struct {
	done bool

	data  []string
	error error

	inputs []Node // Adds could have 0 or 1 input.
	parm   AddParm
}

type AddParm struct {
	adds []string
}

// NewAdd creates a new Add node.
func NewAdd(parm AddParm) *Add {
	n := &Add{
		inputs: make([]Node, 1),
		parm:   parm,
		data:   []string{},
	}
	return n
}

// Type is a type name of the node.
func (n *Add) Type() string {
	return "Add"
}

// Inputs returns it's inputs.
func (n *Add) Inputs() []Node {
	return n.inputs
}

// AddInput set or replaces the first input Node.
func (n *Add) AddInput(in Node) {
	n.inputs[0] = in
}

// Result returns it's calculated data.
func (n *Add) Result() ([]string, error) {
	if !n.done {
		n.add()
	}
	n.done = true

	return n.data, n.error
}

// add adds new data to it's data.
func (n *Add) add() {
	if n.parm.adds == nil {
		n.error = NewError(n, "parm.adds should not nil")
		return
	}

	if n.inputs[0] != nil {
		data, err := n.inputs[0].Result()
		if err != nil {
			n.error = err
			return
		}
		if data == nil {
			n.error = NewError(n, "first input's data should not nil")
			return
		}
		n.data = data
	}

	n.data = append(n.data, n.parm.adds...)
}
