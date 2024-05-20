package query

func Filter[T any](slice []T, f func(T) bool) []T {
	var n []T
	if slice == nil {
		return n
	}
	for _, e := range slice {
		if f(e) {
			n = append(n, e)
		}
	}
	return n
}
