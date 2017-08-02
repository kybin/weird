package stringctx

import (
	"fmt"

	"github.com/kybin/weird/node/stringnode"
)

type Context struct {
	nodes map[string]stringnode.Node
}

// NewContext creates a new Context.
func NewContext() *Context {
	return &Context{
		nodes: make(map[string]stringnode.Node),
	}
}

// ListNodeTypes shows all registered node types in string.
func (c Context) ListNodeTypes() []string {
	return stringnode.List()
}

// Create creates a new stringnode.Node and return it.
// Created node will also registered to Context.nodes.
//
// If the name of already exists in Context, or if the type of node
// is not registered in stringnode, it will return error.
func (c Context) Create(name, typ string) (stringnode.Node, error) {
	if _, ok := c.nodes[name]; ok {
		return nil, fmt.Errorf("node name already exists: %v", name)
	}
	n := stringnode.Create(typ)
	c.node[name] = n
	return n
}
