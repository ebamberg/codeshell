package ext

type Equaliable[T any] interface {
	Equals(other T) bool
}

type Comparable[T any] interface {
	Equaliable[T]
	CompareTo(other T) int
}

type Predicate[T any] func(T) bool

func MatchAll[T any](e T) bool {
	return true
}

type Provider[T any] interface {
	List() []T
	Get(key string) T
}
