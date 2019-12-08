package bptree

import "container/list"

type leaf struct {
	nextLeaf *leaf
	li list.List
}

type leafElement struct {
	key int
	value interface{}
}

func (l *leaf) add (key int, value interface{}) {
	ele := leafElement{key, value}
	for e := l.li.Front(); e != nil; e = e.Next() {
		_ele := e.Value.(leafElement)

		if _ele.key == ele.key {
			e.Value = ele
			return
		}
		if _ele.key > ele.key {
			l.li.InsertBefore(ele, e)
			return
		}
	}
	l.li.PushBack(ele)
}

func (l *leaf) find(key int) (interface{}, bool) {
	for e := l.li.Front(); e != nil; e = e.Next() {
		ele := e.Value.(leafElement)
		if ele.key == key {
			return ele.value, true
		}
	}
	return nil, false
}