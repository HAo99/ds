package stack

import "github.com/HAo99/ds"

type Stack[T any] interface {
	ds.Container[T]

	Push(x T)
	Pop() (T, error)
	Peek() (T, error)
}

type StackX[T any] interface {
	PopX() T
	PeekX() T
}
