package stringnode

type Node interface {
	Name() string
	Type() string
	Inputs() []Node
	AddInput(Node)
	Result() ([]string, error)
}
