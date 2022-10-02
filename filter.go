package fp

func FilterSlice[C any](in []C, pred func(C) bool) []C {
	out := []C{}
	for _, v := range in {
		if pred(v) {
			out = append(out, v)
		}
	}
	return out
}
