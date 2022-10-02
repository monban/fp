package fp

func AllSlice[C any](a []C, pred func(C) bool) bool {
	for _, v := range a {
		if !pred(v) {
			return false
		}
	}
	return true
}
