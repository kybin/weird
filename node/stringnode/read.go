package stringnode

import "io/ioutil"

// ReadNode implements stringnode.Node.
type ReadNode struct {
	done   bool
	result string
	error  error

	parm ReadNodeParm
}

type ReadNodeParm struct {
	fpath string
}

// NewReadNode creates a new ReadNode and initialize it's parameters.
func NewReadNode(parm ReadNodeParm) *ReadNode {
	n := &ReadNode{}
	n.parm = parm
	return n
}

// Inputs implements Node interface, nothing else.
func (n *ReadNode) Inputs() []Node {
	return nil
}

// AddInput will do nothing!
func (n *ReadNode) AddInput(in Node) {}

func (n ReadNode) Result() (string, error) {
	if !n.done {
		n.read()
	}
	n.done = true

	return n.result, n.error
}

// read reads content from a file.
func (n *ReadNode) read() {
	bs, err := ioutil.ReadFile(n.parm.fpath)
	if err != nil {
		n.error = err
		return
	}
	n.result = string(bs)
}
