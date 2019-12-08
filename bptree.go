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
	l := bpt.findLeaf(key)
	l.add(key, value)
}

func (bpt *bptree) find(key int) (interface{}, bool) {
	l := bpt.findLeaf(key)
	return l.find(key)
}

func (bpt *bptree) findLeaf(key int) *leaf {
	if n, ok := bpt.root.(*leaf); ok {
		return n
	}
	return nil
}