package bptree

type branch struct {
	keys []int
	leafs []*leaf
}

func NewBranch(center int, l, r *leaf) *branch {
	b := branch{
		keys:  []int{center},
		leafs: []*leaf{l, r},
	}
	return &b
}

func (b *branch) next(key int) node {
	for i, k := range b.keys {
		if key < k {
			return b.leafs[i]
		}
	}
	return b.leafs[len(b.leafs) - 1]
}
