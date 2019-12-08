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

	bpt.add(3, "b")
	ret, _ = bpt.find(3)
	assert.OK(t, ret == "b")
}