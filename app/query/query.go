package query

import "slices"

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

func RemoveElement[T any](s []T, predicate func(e T) bool) []T {
	if s == nil {
		return make([]T, 0)
	}
	i := slices.IndexFunc(s, predicate)
	if i > -1 {
		return RemoveAtIndex(s, i)
	} else {
		return s
	}
}

func RemoveAtIndex[T any](s []T, index int) []T {
	ret := make([]T, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}
