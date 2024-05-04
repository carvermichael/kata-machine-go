package kata

import (
	. "kata-machine-go/misc"
	"testing"
)

func TestSinglyLinkedList(t *testing.T) {
	list := SinglyLinkedList[int]{}
	TestList(t, &list)
}
