package bptree

import (
	"fmt"
)

var (
	minElementsCount = 2
	maxElementsCount = 4
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

func (bpt *bptree) Add(key int, value interface{}) bool {
	traceBranches := make([]*branch, 0)
	l := bpt.findLeafWithWriteLatch(bpt.root, key, &traceBranches)

	l.rwLatch.Lock()
	defer l.rwLatch.Unlock()
	added, divided, center, newNode := l.add(key, value)

	if !added {
		for _, b := range traceBranches {
			b.rwLatch.Unlock()
		}
		return false
	}

	// propagate to parent branches
	if divided {
		//fmt.Println(traceBranches)
		if len(traceBranches) == 0 {
			b := NewBranch(center, l, newNode.(*leaf))
			bpt.root = b
		} else {
			branchDivided := false
			for i := len(traceBranches)-1; i >= 0; i-- {
				branchDivided, center, newNode = traceBranches[i].add(center, newNode)
				if !branchDivided {
					break
				}
			}
			if branchDivided {
				b := NewBranch(center, traceBranches[0], newNode.(*branch))
				bpt.root = b
			}
		}
	}

	for _, b := range traceBranches {
		b.rwLatch.Unlock()
	}
	return true
}

func (bpt *bptree) Find(key int) (interface{}, bool) {
	l := bpt.findLeaf(bpt.root, key, nil)

	l.rwLatch.RLock()
	defer l.rwLatch.RUnlock()
	return l.find(key)
}

func (bpt *bptree) Delete(key int) {
	traceBranches := make([]*branch, 0)
	l := bpt.findLeafWithWriteLatch(bpt.root, key, &traceBranches)

	l.rwLatch.Lock()
	defer l.rwLatch.Unlock()
	l.delete(key)

	var mergedNode node
	if l != bpt.root && l.wantToMerge() {
		var mergeTarget node = l
		for i := len(traceBranches)-1; i >= 0; i-- {
			mergedNode = traceBranches[i].mergeChildren(mergeTarget)

			// necessary divide mergedNode before merging parent branch if mergedNode want to divide
			switch mergedNode := mergedNode.(type) {
			case *leaf:
				if mergedNode.wantToDivide() {
					center, newLeaf := mergedNode.divide()
					traceBranches[i].add(center, newLeaf)
				}
			case *branch:
				if mergedNode.wantToDivide() {
					center, newNode := mergedNode.divide()
					traceBranches[i].add(center, newNode)
				}
			}

			if !traceBranches[i].wantToMerge() {
				break
			}
			mergeTarget = traceBranches[i]

			if i == 0 && len(traceBranches[i].keys) == 0 {
				// change root if root has no elements
				bpt.root = mergedNode
			}
		}
	}

	for _, b := range traceBranches {
		b.rwLatch.Unlock()
	}
}

func (bpt *bptree) findLeaf(n node, key int, traceBranches *[]*branch) *leaf {
	switch n := n.(type) {
	case *leaf:
		return n
	case *branch:
		n.rwLatch.RLock()
		if traceBranches != nil {
			*traceBranches = append(*traceBranches, n)
		}
		nextNode := n.next(key)
		n.rwLatch.RUnlock()
		return bpt.findLeaf(nextNode, key, traceBranches)
	}
	return nil
}

func (bpt *bptree) findLeafWithWriteLatch(n node, key int, traceBranches *[]*branch) *leaf {
	switch n := n.(type) {
	case *leaf:
		return n
	case *branch:
		if traceBranches != nil {
			*traceBranches = append(*traceBranches, n)
		}
		n.rwLatch.Lock()
		return bpt.findLeafWithWriteLatch(n.next(key), key, traceBranches)
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
