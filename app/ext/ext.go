package ext

type Equaliable[T any] interface {
	Equals(other T) bool
}

type Comparable[T any] interface {
	Equaliable[T]
	CompareTo(other T) int
}
