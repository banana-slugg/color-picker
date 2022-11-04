package util

import "golang.org/x/exp/constraints"

func Max[T constraints.Ordered](vars ...T) T {
	max := vars[0]
	for _, i := range vars {
		if max < i {
			max = i
		}
	}
	return max
}
