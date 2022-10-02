package fp

func MapSlice[C any](a []C, procedure func(C) C) []C {
	b := make([]C, len(a))
	for i, v := range a {
		b[i] = procedure(v)
	}
	return b
}
