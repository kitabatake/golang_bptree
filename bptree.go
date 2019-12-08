package bptree

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
	ok := l.add(key, value)
	if !ok {
		center, newLeaf := l.divide()
		b := NewBranch(center, l, newLeaf)
		newLeaf.add(key, value)
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