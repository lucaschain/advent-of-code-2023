package helpers

import "slices"

type Sorter[T any] func(T, T) int

type Comparator[T any] func(T, T) bool

func SliceAnyEqual[T any](a, b []T, sorter Sorter[T], comparator Comparator[T]) bool {
	if len(a) != len(b) {
		return false
	}

	slices.SortFunc(a, sorter)
	slices.SortFunc(b, sorter)

	for i, v := range a {
		if !comparator(v, b[i]) {
			return false
		}
	}

	return true
}

func SliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	slices.Sort(a)
	slices.Sort(b)

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func SliceSum(a []int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}
