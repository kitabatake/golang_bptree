package bptree

import "fmt"

var (
	m = 3
)

type node interface {
}

type bptree struct {
	root node
}

func NewBptree() *bptree {
	bpt := bptree{}
	root := leaf{}
	bpt.root = &root
	return &bpt
}

func (bpt *bptree) add(key int, value interface{}) {
	l := bpt.findLeaf(bpt.root, key)
	divided, center, newLeaf := l.add(key, value)
	if divided {
		b := NewBranch(center, l, newLeaf.(*leaf))
		//fmt.Printf("divided!\n  center: %d\n  l: %s\n  n: %s\n", center, l, newLeaf)
		bpt.root = b
	}
}

func (bpt *bptree) find(key int) (interface{}, bool) {
	l := bpt.findLeaf(bpt.root, key)
	return l.find(key)
}

func (bpt *bptree) findLeaf(n node, key int) *leaf {
	switch n := n.(type) {
	case *leaf:
		return n
	case *branch:
		return bpt.findLeaf(n.next(key), key)
	}
	return nil
}

func (bpt *bptree) dump() {
	q := []node{bpt.root}
	for len(q) != 0 {
		tq := append([]node{}, q...)
		q = []node{}

		for _, n := range tq {
			fmt.Printf("%s ", n)
			if n, ok := n.(*branch); ok {
				for _, child := range n.nodes {
					q = append(q, child)
				}
			}
		}
		fmt.Println()
	}
}