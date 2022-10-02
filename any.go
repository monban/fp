package fp

// AnySlice returns true if any member of the collection satisfies the predicate
func AnySlice[C any](a []C, pred func(C) bool) bool {
	for _, v := range a {
		if pred(v) {
			return true
		}
	}
	return false
}
