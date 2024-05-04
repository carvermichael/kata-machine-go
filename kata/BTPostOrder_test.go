package kata

import (
	. "github.com/onsi/gomega"
	. "kata-machine-go/misc"
	"testing"
)

func TestBinaryTreePostOrder(t *testing.T) {
	g := NewWithT(t)

	g.Expect(BTPostOrder(Tree)).To(Equal([]int{
		7,
		5,
		15,
		10,
		29,
		45,
		30,
		100,
		50,
		20,
	}))
}
