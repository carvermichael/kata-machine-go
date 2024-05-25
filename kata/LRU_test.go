package kata

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestLRU(t *testing.T) {
	g := NewWithT(t)

	lru := LRU[string, int]{}
	lru.Init(3)

	g.Expect(lru.Get("foo")).Error().To(HaveOccurred())
	lru.Update("foo", 69)
	g.Expect(lru.Get("foo")).To(Equal(69))

	lru.Update("bar", 420)
	g.Expect(lru.Get("bar")).To(Equal(420))

	lru.Update("baz", 1337)
	g.Expect(lru.Get("baz")).To(Equal(1337))

	lru.Update("ball", 69420)
	g.Expect(lru.Get("ball")).To(Equal(69420))
	g.Expect(lru.Get("foo")).Error().To(HaveOccurred())
	g.Expect(lru.Get("bar")).To(Equal(420))
	lru.Update("foo", 69)
	g.Expect(lru.Get("bar")).To(Equal(420))
	g.Expect(lru.Get("foo")).To(Equal(69))

	g.Expect(lru.Get("baz")).Error().To(HaveOccurred())
}
