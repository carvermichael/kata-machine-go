package kata

import (
	. "github.com/onsi/gomega"
	"math/rand"
	"testing"
)

func TestTwoCrystalBalls(t *testing.T) {
	g := NewWithT(t)

	idx := rand.Intn(10000)
	data := make([]bool, 10000)

	for i := idx; i < 10000; i++ {
		data[i] = true
	}

	g.Expect(TwoCrystalBalls(data)).To(Equal(idx))
	g.Expect(TwoCrystalBalls(make([]bool, 10000))).To(Equal(-1))
}
