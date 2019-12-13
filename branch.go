package bptree

import "fmt"

type branch struct {
	keys []int
	nodes []node
}

func NewBranch(center int, l, r node) *branch {
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

func (b *branch) add(key int, n node) (bool, int, node) {
	inserted := false
	for i, k := range b.keys {
		if k > key {
			b.keys = append(b.keys[:i], append([]int{key}, b.keys[i:]...)...)
			b.nodes = append(b.nodes[:i+1], append([]node{n}, b.nodes[i+1:]...)...)
			inserted = true
			break
		}
	}
	if !inserted {
		b.keys = append(b.keys, key)
		b.nodes = append(b.nodes, n)
	}

	if len(b.keys) > maxElementsCount {
		center, newBranch := b.divide()
		return true, center, newBranch
	}

	return false, 0, nil
}

func (b *branch) centerIndex() int {
	m := len(b.keys)
	if m % 2 == 0 {
		return m/2-1
	} else {
		return m/2
	}
}

func (b *branch) divide() (int, node) {
	centerIndex := b.centerIndex()
	center := b.keys[centerIndex]
	newBranch := &branch{}

	newBranch.keys = append([]int{}, b.keys[centerIndex+1:]...)
	newBranch.nodes = append([]node{}, b.nodes[centerIndex+1:]...)
	b.keys = b.keys[:centerIndex]
	b.nodes = b.nodes[:centerIndex+1]

	return center, newBranch
}

func (b *branch) mergeChildren(n node) {
	var deleteKeyIndex, deleteNodeIndex int
	var targetNode node
	for i, _n := range b.nodes {
		if _n == n {
			deleteNodeIndex = i
			if i != len(b.nodes)-1 {
				deleteKeyIndex = i
				targetNode = b.nodes[i+1]
			} else {
				deleteKeyIndex = i-1
				targetNode = b.nodes[i-1]
			}
		}
	}

	if targetLeaf, ok := targetNode.(*leaf); ok {
		targetLeaf.merge(n.(*leaf))
	} else {
		targetNode.(*branch).merge(b.keys[deleteKeyIndex], n.(*branch))
	}
	b.keys = append(b.keys[:deleteKeyIndex], b.keys[deleteKeyIndex+1:]...)
	b.nodes = append(b.nodes[:deleteNodeIndex], b.nodes[deleteNodeIndex+1:]...)
}

func (b *branch) merge(center int, targetBranch *branch) {
	if b.keys[0] > targetBranch.keys[0] {
		b.keys = append(targetBranch.keys, append([]int{center}, b.keys...)...)
		b.nodes = append(targetBranch.nodes, b.nodes...)
	} else {
		b.keys = append(b.keys, append([]int{center}, targetBranch.keys...)...)
		b.nodes = append(b.nodes, targetBranch.nodes...)
	}
}

func (b *branch) wantToMerge() bool {
	return len(b.keys) < minElementsCount
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