package fp

import "golang.org/x/exp/constraints"

// Implements the Quicksort algorythm on a collection of Ordered elements
func QuickSortSlice[C constraints.Ordered](in []C) []C {
	if len(in) < 2 {
		return in
	}
	out := make([]C, len(in))
	var a []C
	var b []C
	pivot := in[0]
	in = in[1:]
	for _, v := range in {
		if v < pivot {
			a = append(a, v)
		} else {
			b = append(b, v)
		}
	}
	a = QuickSortSlice(a)
	b = QuickSortSlice(b)
	out = append(a, pivot)
	out = append(out, b...)
	return out
}
