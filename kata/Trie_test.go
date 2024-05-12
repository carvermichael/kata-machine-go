package kata

import (
	. "github.com/onsi/gomega"
	"sort"
	"testing"
)

func TestTrie(t *testing.T) {
	g := NewWithT(t)

	trie := Trie{}

	trie.Insert("foo")
	trie.Insert("fool")
	trie.Insert("foolish")
	trie.Insert("bar")
	trie.Insert("foobar")

	findResult := trie.Find("fo")
	sort.Strings(findResult)

	g.Expect(findResult).To(Equal([]string{
		"foo",
		"foobar",
		"fool",
		"foolish",
	}))

	trie.Delete("fool")

	findResult = trie.Find("fo")
	sort.Strings(findResult)

	g.Expect(findResult).To(Equal([]string{
		"foo",
		"foobar",
		"foolish",
	}))

	trie.Delete("foobar")

	findResult = trie.Find("fo")
	sort.Strings(findResult)

	g.Expect(findResult).To(Equal([]string{
		"foo",
		"foolish",
	}))
}
