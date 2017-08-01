package stringnode

type Node interface {
	Type() string
	Inputs() []Node
	AddInput(Node)
	Result() ([]string, error)
}
