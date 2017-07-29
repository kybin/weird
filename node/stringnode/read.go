package stringnode

import "io/ioutil"

// Read implements stringnode.Node.
type Read struct {
	done   bool
	result []string
	error  error

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

	return n.result, n.error
}

// VeiledResult implements node.Node.
func (n *Read) VeiledResult() (interface{}, error) {
	data, err := n.Result()
	return interface{}(data), err
}

// read reads content from a file.
func (n *Read) read() {
	bs, err := ioutil.ReadFile(n.parm.fpath)
	if err != nil {
		n.error = err
		return
	}
	n.result = []string{string(bs)}
}
