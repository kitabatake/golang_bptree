package bptree

import (
	"container/list"
	"fmt"
)

type leaf struct {
	nextLeaf *leaf
	li list.List
}

type leafElement struct {
	key int
	value interface{}
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

func (l *leaf) add (key int, value interface{}) (bool, int, node) {
	if l.update(key, value) {
		return false, 0, nil
	}

	if l.li.Len() >= m {
		center, newLeaf := l.divide(key, value)
		return true, center, newLeaf
	}

	ele := leafElement{key, value}
	for e := l.li.Front(); e != nil; e = e.Next() {
		_ele := e.Value.(leafElement)
		if _ele.key > ele.key {
			l.li.InsertBefore(ele, e)
			return false, 0, nil
		}
	}
	l.li.PushBack(ele)
	return false, 0, nil
}

func (l *leaf) update(key int, value interface{}) bool {
	for e := l.li.Front(); e != nil; e = e.Next() {
		ele := e.Value.(leafElement)

		if ele.key == key {
			e.Value = leafElement{key, value}
			return true
		}
	}
	return false
}

func (l *leaf) divide(newKey int, newValue interface{}) (int, *leaf) {
	keys := make([]int, 0)
	values := make([]interface{}, 0)
	added := false
	for e := l.li.Front(); e != nil; e = e.Next() {
		ele := e.Value.(leafElement)
		if !added && ele.key > newKey {
			keys = append(keys, newKey)
			values = append(values, newValue)
			added = true
		}
		keys = append(keys, ele.key)
		values = append(values, ele.value)
	}

	if !added {
		keys = append(keys, newKey)
		values = append(values, newValue)
	}

	centerIndex := l.centerIndex()
	l.li.Init()
	for i, k := range keys[:centerIndex] {
		l.li.PushBack(leafElement{k, values[i]})
	}

	newLeaf := &leaf{}
	for i, k := range keys[centerIndex:] {
		newLeaf.li.PushBack(leafElement{k, values[centerIndex+i]})
	}

	l.nextLeaf = newLeaf
	return keys[centerIndex], newLeaf
}


func (l *leaf) String() string {
	out := "["
	listLen := l.li.Len()
	i := 0
	for e := l.li.Front(); e != nil; e = e.Next() {
		ele := e.Value.(leafElement)
		out += fmt.Sprintf("%d(%s)", ele.key, ele.value)
		if i != (listLen -1) {
			out += ", "
		}
		i++
	}
	return out + "]"
}