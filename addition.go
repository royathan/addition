package addition

import (
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

// Adds two integers together via [addition].
//
// [addition]: https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](a, b T) T {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

