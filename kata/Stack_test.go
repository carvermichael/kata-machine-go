package kata

import (
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
	"testing"
)

func TestStack(t *testing.T) {
	g := NewWithT(t)

	stack := Stack[int]{}

	stack.Push(5)
	stack.Push(7)
	stack.Push(9)

	g.Expect(stack.Pop()).To(PointTo(Equal(9)))
	g.Expect(stack.Length()).To(Equal(2))

	stack.Push(11)
	g.Expect(stack.Pop()).To(PointTo(Equal(11)))
	g.Expect(stack.Pop()).To(PointTo(Equal(7)))
	g.Expect(stack.Peek()).To(PointTo(Equal(5)))
	g.Expect(stack.Pop()).To(PointTo(Equal(5)))
	g.Expect(stack.Pop()).Error().To(HaveOccurred())

	stack.Push(69)
	g.Expect(stack.Peek()).To(PointTo(Equal(69)))
	g.Expect(stack.Length()).To(Equal(1))
}
