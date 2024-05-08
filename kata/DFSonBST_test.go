package kata

import (
	. "github.com/onsi/gomega"
	. "kata-machine-go/misc"
	"testing"
)

func TestDFSonBST(t *testing.T) {
	g := NewWithT(t)

	g.Expect(DFSonBST(Tree, 45)).To(Equal(true))
	g.Expect(DFSonBST(Tree, 7)).To(Equal(true))
	g.Expect(DFSonBST(Tree, 69)).To(Equal(false))
}
