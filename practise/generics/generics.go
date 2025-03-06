package generics

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// 1. generic function

func Print[T any](v T) {
	fmt.Println(v)
}

// 2. constraint generic function
func SumMap[K comparable, V int | float64](m map[K]V) V {
	var sum V
	for _, v := range m {
		sum += v
	}
	return sum
}

func Add[T constraints.Ordered](a, b T) T {
	return a + b
}

// 3. custome type
type Number interface {
	~int | ~float32 | ~float64
}

func AddNumber[T Number](a, b T) T {
	return a + b
}

// 4. generic struct
type Container[T any] struct {
	Value T
}

func NewContainer[T any](v T) Container[T] {
	return Container[T]{Value: v}
}

// 5. generic method
func (c *Container[T]) SetVal(v T) {
	c.Value = v
}

// 6. generic slice
func Map[T any, R any](arr []T, fn func(T) R) []R {
	var res []R
	for _, v := range arr {
		res = append(res, fn(v))
	}
	return res
}
