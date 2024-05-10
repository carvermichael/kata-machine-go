package kata

import (
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
	"testing"
)

func TestMinHeap(t *testing.T) {
	g := NewWithT(t)

	heap := MinHeap[int]{}

	heap.Insert(5)
	heap.Insert(3)
	heap.Insert(69)
	heap.Insert(420)
	heap.Insert(4)
	heap.Insert(1)
	heap.Insert(8)
	heap.Insert(7)

	g.Expect(heap.Length()).To(Equal(8))
	g.Expect(heap.Delete()).To(PointTo(Equal(1)))
	g.Expect(heap.Delete()).To(PointTo(Equal(3)))
	g.Expect(heap.Delete()).To(PointTo(Equal(4)))
	g.Expect(heap.Delete()).To(PointTo(Equal(5)))
	g.Expect(heap.Length()).To(Equal(4))
	g.Expect(heap.Delete()).To(PointTo(Equal(7)))
	g.Expect(heap.Delete()).To(PointTo(Equal(8)))
	g.Expect(heap.Delete()).To(PointTo(Equal(69)))
	g.Expect(heap.Delete()).To(PointTo(Equal(420)))
	g.Expect(heap.Length()).To(Equal(0))
	g.Expect(heap.Delete()).Error().To(HaveOccurred())
}
