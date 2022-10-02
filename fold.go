package fp

func FoldSlice[C any](s []C, reduction func(C, C) C) C {
	var out C
	for _, v := range s {
		out = reduction(out, v)
	}
	return out
}
