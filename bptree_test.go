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

	_, ok := bpt.find(1)
	assert.OK(t, ok == false)

	bpt.add(1, "a")
	ret, _ := bpt.find(1)
	assert.OK(t, ret == "a")

	bpt.add(1, "aa")
	ret, _ = bpt.find(1)
	assert.OK(t, ret == "aa")

	bpt.add(3, "b")
	ret, _ = bpt.find(3)
	assert.OK(t, ret == "b")

	bpt.add(4, "c")
	ret, _ = bpt.find(4)
	assert.OK(t, ret == "c")

	bpt.add(5, "d")
	ret, _ = bpt.find(5)
	//bpt.dump()
	assert.OK(t, ret == "d")
}

func TestExpansion(t *testing.T) {
	bpt := NewBptree()
	for i := 1; i <= 10; i++ {
		bpt.add(i, i)
	}
	//bpt.dump()
	for i := 1; i <= 10; i++ {
		ret, _ := bpt.find(i)
		assert.OK(t, ret == i)
	}
}

func TestDescendingOrderExpansion(t *testing.T) {
	bpt := NewBptree()
	for i := 10; i >= 1; i-- {
		bpt.add(i, i)
	}
	bpt.dump()
	for i := 10; i >= 1; i-- {
		ret, _ := bpt.find(i)
		assert.OK(t, ret == i)
	}
}