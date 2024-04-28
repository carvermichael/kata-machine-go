package kata

import (
	. "kata-machine-go/testUtils"
	"testing"
)

func TestSinglyLinkedList(t *testing.T) {
	list := SinglyLinkedList[int]{}
	TestList(t, &list)
}
