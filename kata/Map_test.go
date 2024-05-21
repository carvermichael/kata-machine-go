package kata

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestMap(t *testing.T) {
	g := NewWithT(t)

	mappy := Map{}
	mappy.Init()

	mappy.Set("foo", 55)
	g.Expect(mappy.Size()).To(Equal(1))
	mappy.Set("fool", 75)
	g.Expect(mappy.Size()).To(Equal(2))
	mappy.Set("foolish", 105)
	g.Expect(mappy.Size()).To(Equal(3))
	mappy.Set("bar", 69)
	g.Expect(mappy.Size()).To(Equal(4))

	g.Expect(mappy.Get("bar")).To(Equal(69))
	g.Expect(mappy.Get("blaz")).Error().To(HaveOccurred())

	mappy.Delete("barblabr")
	g.Expect(mappy.Size()).To(Equal(4))

	mappy.Delete("bar")
	g.Expect(mappy.Size()).To(Equal(3))
	g.Expect(mappy.Get("bar")).Error().To(HaveOccurred())
}
