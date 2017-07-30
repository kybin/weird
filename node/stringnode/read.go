package stringnode

import "io/ioutil"

// Read implements stringnode.Node.
type Read struct {
	done bool

	data  []string
	error error

	name string
	parm ReadParm
}

type ReadParm struct {
	fpath string
}

// NewRead creates a new Read and initialize it's parameters.
func NewRead(name string, parm ReadParm) *Read {
	n := &Read{
		name: name,
		parm: parm,
	}
	return n
}

// Name is a name of the node.
func (n *Read) Name() string {
	return n.name
}

// Type is a type name of the node.
func (n *Read) Type() string {
	return "Read"
}

// Inputs implements Node interface, nothing else.
func (n *Read) Inputs() []Node {
	return nil
}

// AddInput will do nothing!
func (n *Read) AddInput(in Node) {}

func (n Read) Result() ([]string, error) {
	if !n.done {
		n.read()
	}
	n.done = true

	return n.data, n.error
}

// read reads content from a file.
func (n *Read) read() {
	bs, err := ioutil.ReadFile(n.parm.fpath)
	if err != nil {
		n.error = NewError(n, err.Error())
		return
	}
	n.data = []string{string(bs)}
}
