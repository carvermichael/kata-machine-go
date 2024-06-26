package kata

import (
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
	"testing"
)

func TestQueue(t *testing.T) {
	g := NewWithT(t)

	queue := Queue[int]{}

	queue.Enqueue(5)
	queue.Enqueue(7)
	queue.Enqueue(9)

	g.Expect(queue.Dequeue()).To(PointTo(Equal(5)))
	g.Expect(queue.Length()).To(Equal(2))

	queue.Enqueue(11)
	g.Expect(queue.Dequeue()).To(PointTo(Equal(7)))
	g.Expect(queue.Dequeue()).To(PointTo(Equal(9)))
	g.Expect(queue.Peek()).To(PointTo(Equal(11)))
	g.Expect(queue.Dequeue()).To(PointTo(Equal(11)))
	g.Expect(queue.Dequeue()).Error().To(HaveOccurred())
	g.Expect(queue.Length()).To(Equal(0))

	queue.Enqueue(69)
	g.Expect(queue.Peek()).To(PointTo(Equal(69)))
	g.Expect(queue.Length()).To(Equal(1))
}
