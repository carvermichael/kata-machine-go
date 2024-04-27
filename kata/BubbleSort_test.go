package kata

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	g := NewWithT(t)

	arr := []int{9, 3, 7, 4, 69, 420, 42}

	BubbleSort(&arr)
	g.Expect(arr).To(Equal([]int{3, 4, 7, 9, 42, 69, 420}))
}
