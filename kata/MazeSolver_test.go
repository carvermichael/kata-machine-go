package kata

import (
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
	. "kata-machine-go/types"
	"testing"
)

func TestMazeSolver(t *testing.T) {
	g := NewGomegaWithT(t)

	maze := []string{
		"xxxxxxxxxx x",
		"x        x x",
		"x        x x",
		"x xxxxxxxx x",
		"x          x",
		"x xxxxxxxxxx",
	}

	expected := []Point{
		{X: 10, Y: 0},
		{X: 10, Y: 1},
		{X: 10, Y: 2},
		{X: 10, Y: 3},
		{X: 10, Y: 4},
		{X: 9, Y: 4},
		{X: 8, Y: 4},
		{X: 7, Y: 4},
		{X: 6, Y: 4},
		{X: 5, Y: 4},
		{X: 4, Y: 4},
		{X: 3, Y: 4},
		{X: 2, Y: 4},
		{X: 1, Y: 4},
		{X: 1, Y: 5},
	}

	actual := MazeSolver(maze, 'x', Point{X: 10, Y: 0}, Point{X: 1, Y: 5})

	g.Expect(actual).To(PointTo(Equal(expected)))
}
