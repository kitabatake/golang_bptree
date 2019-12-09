package bptree

import "fmt"

type branch struct {
	keys []int
	nodes []node
}

func NewBranch(center int, l, r *leaf) *branch {
	b := branch{
		keys:  []int{center},
		nodes: []node{l, r},
	}
	return &b
}

func (b *branch) next(key int) node {
	for i, k := range b.keys {
		if key < k {
			return b.nodes[i]
		}
	}
	return b.nodes[len(b.nodes) - 1]
}

func (b *branch) String() string {
	out := "["
	for i, k := range b.keys {
		out += fmt.Sprintf("%d", k)
		if i < len(b.keys)-1 {
			out += ", "
		}
	}
	return out + "]"
}