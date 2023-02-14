package min

import "golang.org/x/exp/constraints"

func Min[A constraints.Ordered](x, y A) A {
	if x <= y {
		return x
	}
	return y
}

func Max[A constraints.Ordered](x, y A) A {
	if x >= y {
		return x
	}
	return y
}
