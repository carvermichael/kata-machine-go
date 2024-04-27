package kata

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	g := NewWithT(t)

	haystack := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}

	g.Expect(LinearSearch(haystack, 69), true)
	g.Expect(LinearSearch(haystack, 1336), false)
	g.Expect(LinearSearch(haystack, 69420), true)
	g.Expect(LinearSearch(haystack, 69421), false)
	g.Expect(LinearSearch(haystack, 1), true)
	g.Expect(LinearSearch(haystack, 0), false)
}
