package bptree

import "fmt"

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

func (bpt *bptree) Add(key int, value interface{}) {
	traceBranches := make([]*branch, 0)
	l := bpt.findLeaf(bpt.root, key, &traceBranches)
	divided, center, newNode := l.add(key, value)

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
}

func (bpt *bptree) Find(key int) (interface{}, bool) {
	l := bpt.findLeaf(bpt.root, key, nil)
	return l.find(key)
}

func (bpt *bptree) Delete(key int) {
	traceBranches := make([]*branch, 0)
	l := bpt.findLeaf(bpt.root, key, &traceBranches)
	l.delete(key)

	var mergedNode node
	if l != bpt.root && l.wantToMerge() {
		var mergeTarget node = l
		for i := len(traceBranches)-1; i >= 0; i-- {
			//fmt.Printf("merge! %s\n", mergeTarget)
			mergedNode = traceBranches[i].mergeChildren(mergeTarget)
			if !traceBranches[i].wantToMerge() {
				break
			}
			mergeTarget = traceBranches[i]

			if i == 0 && len(traceBranches[i].keys) == 0 {
				// change root is root has no elements
				bpt.root = mergedNode
			}
		}
	}
}

func (bpt *bptree) findLeaf(n node, key int, traceBranches *[]*branch) *leaf {
	switch n := n.(type) {
	case *leaf:
		return n
	case *branch:
		if traceBranches != nil {
			*traceBranches = append(*traceBranches, n)
		}
		return bpt.findLeaf(n.next(key), key, traceBranches)
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