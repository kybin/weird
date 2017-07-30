package stringnode

import "fmt"

type Error struct {
	Node  Node
	Value string
}

func (e *Error) Error() string {
	return fmt.Sprintf("%v (%v): %v", e.Node.Name(), e.Node.Type(), e.Value)
}

func NewError(n Node, v string) *Error {
	return &Error{
		Node:  n,
		Value: v,
	}
}
