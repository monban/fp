package fp

// AllSlice returns a boolean representing whether the predicate returns true for all members of the collection.
func AllSlice[C any](a []C, pred func(C) bool) bool {
	for _, v := range a {
		if !pred(v) {
			return false
		}
	}
	return true
}
