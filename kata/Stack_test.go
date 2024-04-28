package kata

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestStack(t *testing.T) {
	g := NewWithT(t)

	stack := Stack[int]{}

	stack.Push(5)
	stack.Push(7)
	stack.Push(9)

	g.Expect(stack.Pop()).To(Equal(9))
	g.Expect(stack.Length()).To(Equal(2))

	stack.Push(11)
	g.Expect(stack.Pop()).To(Equal(11))
	g.Expect(stack.Pop()).To(Equal(7))
	g.Expect(stack.Peek()).To(Equal(5))
	g.Expect(stack.Pop()).To(Equal(5))
	g.Expect(stack.Pop()).Error().To(HaveOccurred())

	stack.Push(69)
	g.Expect(stack.Peek()).To(Equal(69))
	g.Expect(stack.Length()).To(Equal(1))
}
