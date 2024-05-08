package kata

import (
	. "github.com/onsi/gomega"
	. "kata-machine-go/misc"
	"testing"
)

func TestBinaryTreeBreadthFirstSearch(t *testing.T) {
	g := NewWithT(t)

	g.Expect(BTBFS(Tree, 45)).To(Equal(true))
	g.Expect(BTBFS(Tree, 7)).To(Equal(true))
	g.Expect(BTBFS(Tree, 69)).To(Equal(false))
}
