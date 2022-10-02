package fp

// FilterSlice creates a new slice of all elements for which the predicate returns false.
func FilterSlice[C any](in []C, pred func(C) bool) []C {
	out := []C{}
	for _, v := range in {
		if pred(v) {
			out = append(out, v)
		}
	}
	return out
}
