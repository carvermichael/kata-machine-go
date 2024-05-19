package kata

import (
	. "github.com/onsi/gomega"
	. "kata-machine-go/misc"
	"testing"
)

func TestBFSGraphList(t *testing.T) {
	g := NewWithT(t)

	g.Expect(BFSGraphList(List2, 0, 6)).To(Equal([]int{
		0,
		1,
		4,
		5,
		6,
	}))

	g.Expect(BFSGraphList(List2, 6, 0)).To(Equal(make([]int, 0)))
}
