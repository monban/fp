package fp

func Any[C any](a []C, pred func(C) bool) bool {
	for _, v := range a {
		if pred(v) {
			return true
		}
	}
	return false
}