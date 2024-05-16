package kata

import (
	. "github.com/onsi/gomega"
	. "kata-machine-go/misc"
	"testing"
)

func TestBFSGraphMatrix(t *testing.T) {
	g := NewWithT(t)

	g.Expect(BFS(Matrix2, 0, 6)).To(Equal([]int{
		0,
		1,
		4,
		5,
		6,
	}))

	g.Expect(BFS(Matrix2, 6, 0)).To(Equal(make([]int, 0)))
}
