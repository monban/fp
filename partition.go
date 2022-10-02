package fp

func Partition[C any](col []C, pred func(C) bool) ([]C, []C) {
	var a []C
	var b []C

	for _, v := range col {
		if pred(v) {
			a = append(a, v)
		} else {
			b = append(b, v)
		}
	}
	return a, b
}
