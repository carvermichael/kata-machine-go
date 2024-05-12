package misc

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestList(t *testing.T, list List[int]) {
	g := NewWithT(t)

	list.Append(5)
	list.Append(7)
	list.Append(9)

	g.Expect(list.Get(2)).To(Equal(7))
	g.Expect(list.RemoveAt(1)).To(Equal(7))
	g.Expect(list.Length()).To(Equal(2))

	list.Append(11)
	g.Expect(list.RemoveAt(1)).To(Equal(9))
	g.Expect(list.Remove(9)).To(Equal(false))
	g.Expect(list.RemoveAt(0)).To(Equal(5))
	g.Expect(list.RemoveAt(0)).To(Equal(11))
	g.Expect(list.Length()).To(Equal(0))

	list.Prepend(5)
	list.Prepend(7)
	list.Prepend(9)

	g.Expect(list.Get(2)).To(Equal(5))
	g.Expect(list.Get(0)).To(Equal(9))
	g.Expect(list.Remove(9)).To(Equal(true))
	g.Expect(list.Length()).To(Equal(2))
	g.Expect(list.Get(0)).To(Equal(7))
}
