package fp

// PartitionSlice returns two arrays, the first containing elements which satify the predicate, the second containing elements which do not.
func PartitionSlice[C any](col []C, pred func(C) bool) ([]C, []C) {
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
