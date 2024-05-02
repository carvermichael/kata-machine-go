package kata

import (
	. "kata-machine-go/testUtils"
	"testing"
)

func TestArrayList(t *testing.T) {
	list := ArrayList[int]{}
	TestList(t, &list)
}
