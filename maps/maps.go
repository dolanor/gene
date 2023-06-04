package maps

func Clone[T map[K]V, K comparable, V any](m map[K]V) map[K]V {
	clone := map[K]V{}
	for k, v := range m {
		clone[k] = v
	}
	return clone
}
