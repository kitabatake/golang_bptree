package bptree

type leaf struct {
	nextLeaf *leaf
	keys []int
	values []interface{}
}

func (l *leaf) add (key int, value interface{}) {
	l.keys = append(l.keys, key)
	l.values = append(l.values, value)
}
