package stringnode

type Node interface {
	Inputs() []Node
	AddInput(Node)
	Result() ([]string, error)
}
