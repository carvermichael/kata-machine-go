package kata

import (
	. "github.com/onsi/gomega"
	. "kata-machine-go/misc"
	"testing"
)

func TestBinaryTreeInOrder(t *testing.T) {
	g := NewWithT(t)

	g.Expect(BTInOrder(Tree)).To(Equal([]int{
		5,
		7,
		10,
		15,
		20,
		29,
		30,
		45,
		50,
		100,
	}))
}
