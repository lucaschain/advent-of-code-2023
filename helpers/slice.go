package helpers

import "slices"

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
