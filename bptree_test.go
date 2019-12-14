package bptree

import (
	"flag"
	"github.com/ToQoz/gopwt"
	"github.com/ToQoz/gopwt/assert"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	flag.Parse()
	gopwt.Empower()
	os.Exit(m.Run())
}

func TestCommon(t *testing.T) {
	bpt := NewBptree()

	_, ok := bpt.Find(1)
	assert.OK(t, ok == false)

	bpt.Add(1, "a")
	ret, _ := bpt.Find(1)
	assert.OK(t, ret == "a")

	bpt.Add(1, "aa")
	ret, _ = bpt.Find(1)
	assert.OK(t, ret == "aa")

	bpt.Add(3, "b")
	ret, _ = bpt.Find(3)
	assert.OK(t, ret == "b")

	bpt.Add(4, "c")
	ret, _ = bpt.Find(4)
	assert.OK(t, ret == "c")

	bpt.Add(5, "d")
	ret, _ = bpt.Find(5)
	//bpt.dump()
	assert.OK(t, ret == "d")

	checkBPTreeCondition(bpt, t)
}

func TestExpansion(t *testing.T) {
	bpt := NewBptree()
	for i := 1; i <= 10; i++ {
		bpt.Add(i, i)
	}
	//bpt.dump()
	for i := 1; i <= 10; i++ {
		ret, _ := bpt.Find(i)
		assert.OK(t, ret == i)
	}

	checkBPTreeCondition(bpt, t)
}

func TestDescendingOrderExpansion(t *testing.T) {
	bpt := NewBptree()
	for i := 10; i >= 1; i-- {
		bpt.Add(i, i)
	}
	for i := 10; i >= 1; i-- {
		ret, _ := bpt.Find(i)
		assert.OK(t, ret == i)
	}

	checkBPTreeCondition(bpt, t)
}

func TestDeleteLeafElement(t *testing.T) {
	bpt := NewBptree()
	bpt.Add(1,1)
	bpt.Add(2,2)
	bpt.Add(4,4)
	bpt.Add(5,4)

	bpt.Delete(4)
	_, ok := bpt.Find(4)
	assert.OK(t, ok == false)

	checkBPTreeCondition(bpt, t)
}

func TestDeleteAndMergeLeafs(t *testing.T) {
	bpt := NewBptree()
	for i := 8; i >= 1; i-- {
		bpt.Add(i, i)
	}


	bpt.Delete(1)

	checkBPTreeCondition(bpt, t)
}

func TestDeleteAndMergeBranches(t *testing.T) {
	bpt := NewBptree()
	for i := 21; i >= 1; i-- {
		bpt.Add(i, i)
	}

	bpt.dump()
	bpt.Delete(21)
	bpt.Delete(20)
	bpt.dump()

	for i := 19; i >= 1; i-- {
		ret, _ := bpt.Find(i)
		assert.OK(t, ret == i)
	}

	checkBPTreeCondition(bpt, t)
}

/**
 * check follows conditions
 * - each elements length between min and max
 * - all leaf nodes level is same
 */
func checkBPTreeCondition(bpt *bptree, t *testing.T) {
	qu := []node{bpt.root}
	var isLeafLevel bool
	for len(qu) != 0 {
		nextQu := make([]node, 0)
		isLeafLevel = false
		for i, n := range qu {
			switch n := n.(type) {
			case *branch:
				assert.OK(t, isLeafLevel == false)

				elementsLen := len(n.keys)
				if n != bpt.root {
					assert.OK(t, elementsLen >= minElementsCount)
				}
				assert.OK(t, elementsLen <= maxElementsCount)
				nextQu = append(nextQu, n.nodes...)
			case *leaf:
				if i == 0 {
					isLeafLevel = true
				} else {
					assert.OK(t, isLeafLevel == true)
				}

				elementsLen := n.li.Len()
				if n != bpt.root {
					assert.OK(t, elementsLen >= minElementsCount)
				}
				assert.OK(t, elementsLen <= maxElementsCount)
			}
		}
		qu = nextQu
	}
}