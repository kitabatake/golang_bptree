package bptree

import (
	"container/list"
	"fmt"
	"sync"
)

type leaf struct {
	nextLeaf *leaf
	li list.List
	rwLatch sync.RWMutex
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
	m := l.li.Len()
	if m % 2 == 0 {
		return m/2
	} else {
		return m/2 + 1
	}
}

func (l *leaf) add (key int, value interface{}) (bool, bool, int, node) {
	if _, ok := l.find(key); ok {
		return false, false, 0, nil
	}

	ele := leafElement{key, value}
	inserted := false
	for e := l.li.Front(); e != nil; e = e.Next() {
		_ele := e.Value.(leafElement)
		if _ele.key > ele.key {
			l.li.InsertBefore(ele, e)
			inserted = true
			break
		}
	}
	if !inserted {
		l.li.PushBack(ele)
	}

	if l.wantToDivide() {
		center, newLeaf := l.divide()
		return true, true, center, newLeaf
	}

	return true, false, 0, nil
}

func (l *leaf) wantToDivide() bool {
	return l.li.Len() > maxElementsCount
}

func (l *leaf) divide() (int, *leaf) {
	keys := make([]int, 0)
	values := make([]interface{}, 0)
	for e := l.li.Front(); e != nil; e = e.Next() {
		ele := e.Value.(leafElement)
		keys = append(keys, ele.key)
		values = append(values, ele.value)
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

func (l *leaf) delete(key int) {
	for e := l.li.Front(); e != nil; e = e.Next() {
		ele := e.Value.(leafElement)
		if ele.key == key {
			l.li.Remove(e)
			break
		}
	}
}

func (l *leaf) wantToMerge() bool {
	return l.li.Len() < minElementsCount
}

func (l *leaf) merge(targetLeaf *leaf) {
	myFirstElement := l.li.Front().Value.(leafElement)
	targetFirstElement := targetLeaf.li.Front().Value.(leafElement)
	if myFirstElement.key > targetFirstElement.key {
		l.li.PushFrontList(&targetLeaf.li)
		l.nextLeaf = targetLeaf.nextLeaf
	} else {
		l.li.PushBackList(&targetLeaf.li)
	}
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