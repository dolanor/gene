package slices

func Clone[T comparable](tt []T) []T {
	clone := make([]T, 0, len(tt))
	for _, v := range tt {
		clone = append(clone, v)
	}
	return clone
}
