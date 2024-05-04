package kata

import (
	. "github.com/onsi/gomega"
	. "kata-machine-go/misc"
	"testing"
)

func TestBinaryTreePreOrder(t *testing.T) {
	g := NewWithT(t)

	g.Expect(BTPreOrder(Tree)).To(Equal([]int{
		20,
		10,
		5,
		7,
		15,
		50,
		30,
		29,
		45,
		100,
	}))
}
