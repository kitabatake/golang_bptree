package bptree

import (
	"container/list"
)

type leaf struct {
	nextLeaf *leaf
	li list.List
}

type leafElement struct {
	key int
	value interface{}
}

func (l *leaf) add (key int, value interface{}) bool {
	if l.li.Len() >= m {
		return false
	}

	ele := leafElement{key, value}
	for e := l.li.Front(); e != nil; e = e.Next() {
		_ele := e.Value.(leafElement)

		if _ele.key == ele.key {
			e.Value = ele
			return true
		}
		if _ele.key > ele.key {
			l.li.InsertBefore(ele, e)
			return true
		}
	}
	l.li.PushBack(ele)
	return true
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

func (l *leaf) centerIndex() int {
	if m % 2 == 0 {
		return m/2
	} else {
		return m/2 + 1
	}
}

func (l *leaf) divide() (int, *leaf) {
	var center int
	centerIndex := l.centerIndex()

	newLeaf := &leaf{}
	i := 0
	for e := l.li.Front(); e != nil; e = e.Next() {
		ele := e.Value.(leafElement)
		if i == centerIndex {
			center = ele.key
		}
		if i >= centerIndex {
			newLeaf.li.PushBack(ele)
			l.li.Remove(e)
		}
		i++
	}

	l.nextLeaf = newLeaf
	return center, newLeaf
}