package kata

import (
	. "github.com/onsi/gomega"
	. "kata-machine-go/misc"
	"testing"
)

func TestDFSonBST(t *testing.T) {
	g := NewWithT(t)

	g.Expect(DFSOnBST(Tree, 45)).To(Equal(true))
	g.Expect(DFSOnBST(Tree, 7)).To(Equal(true))
	g.Expect(DFSOnBST(Tree, 69)).To(Equal(false))
}
