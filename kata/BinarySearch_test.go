package kata

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestBinarySearchList(t *testing.T) {
	g := NewWithT(t)

	haystack := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}

	g.Expect(BinarySearchList(haystack, 69)).To(Equal(true))
	g.Expect(BinarySearchList(haystack, 1336)).To(Equal(false))
	g.Expect(BinarySearchList(haystack, 69420)).To(Equal(true))
	g.Expect(BinarySearchList(haystack, 69421)).To(Equal(false))
	g.Expect(BinarySearchList(haystack, 1)).To(Equal(true))
	g.Expect(BinarySearchList(haystack, 0)).To(Equal(false))
}
