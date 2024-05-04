package kata

import (
	. "kata-machine-go/misc"
	"testing"
)

func TestArrayList(t *testing.T) {
	list := ArrayList[int]{}
	TestList(t, &list)
}
