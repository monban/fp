package fp

// MapSlice returns a collection the same size as the input collection, where each output is the result of applying procedure to the input element.
func MapSlice[C any](a []C, procedure func(C) C) []C {
	b := make([]C, len(a))
	for i, v := range a {
		b[i] = procedure(v)
	}
	return b
}
