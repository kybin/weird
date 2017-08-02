package stringnode

import (
	"fmt"
	"sort"
)

// register nodes here.
var createFn = map[string]func() Node{
	"Add":     func() Node { return NewAdd(AddParm{}) },
	"Join":    func() Node { return NewJoin(JoinParm{}) },
	"Read":    func() Node { return NewRead(ReadParm{}) },
	"Replace": func() Node { return NewReplace(ReplaceParm{n: -1}) },
}

// List is registred node types in string.
func List() []string {
	nodes := []string{}
	for n := range createFn {
		nodes = append(nodes, n)
	}
	sort.Strings(nodes)
	return nodes
}

// Create creates a Node from it's type name.
// If the name is unknown, it will return nil Node.
func Create(name string) (Node, error) {
	create, ok := createFn[name]
	if !ok {
		return nil, fmt.Errorf("unknown node type: %v", name)
	}
	return create(), nil
}
