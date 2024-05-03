package kata

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestQuickSort(t *testing.T) {
	g := NewGomegaWithT(t)

	arr := &[]int{9, 3, 7, 4, 69, 420, 42}

	QuickSort(arr)
	g.Expect(*arr).To(Equal([]int{3, 4, 7, 9, 42, 69, 420}))
}
