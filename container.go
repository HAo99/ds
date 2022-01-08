package ds

type Container[T any] interface {
	Len() int
	Empty() bool
}
