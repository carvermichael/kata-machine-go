package kata

import (
	. "github.com/onsi/gomega"
	. "kata-machine-go/misc"
	"testing"
)

func TestCompareBinaryTrees(t *testing.T) {
	g := NewWithT(t)

	g.Expect(CompareBinaryTrees(Tree, Tree2)).To(Equal(false))
	g.Expect(CompareBinaryTrees(Tree, Tree)).To(Equal(true))
	g.Expect(CompareBinaryTrees(Tree2, Tree2)).To(Equal(true))
}
