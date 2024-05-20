package kata

import (
	. "github.com/onsi/gomega"
	. "kata-machine-go/misc"
	"testing"
)

func TestDijkstraList(t *testing.T) {
	g := NewWithT(t)

	g.Expect(DijkstraList(List2, 0, 6)).To(Equal([]int{
		0,
		1,
		4,
		5,
		6,
	}))
}
