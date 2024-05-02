package kata

import (
	. "github.com/onsi/gomega"
	"math/rand"
	"testing"
)

// Case: Basic random target case.
func TestTwoCrystalBalls_Basic(t *testing.T) {
	g := NewWithT(t)
	WithNumbers(g, 10000, rand.Intn(10000))
}

// Case: the first ball finds the correct floor (w/ sqrt strategy).
func TestTwoCrystalBalls_FirstBall(t *testing.T) {
	g := NewWithT(t)
	WithNumbers(g, 7491, 7224)
}

// Case: Basic random target case.
func TestTwoCrystalBalls_NoBreak(t *testing.T) {
	g := NewWithT(t)
	g.Expect(TwoCrystalBalls(make([]bool, 10000))).To(Equal(-1))
}

func WithNumbers(g *GomegaWithT, size int, target int) {

	data := make([]bool, size)

	for i := target; i < size; i++ {
		data[i] = true
	}

	g.Expect(TwoCrystalBalls(data)).To(Equal(target))
}
