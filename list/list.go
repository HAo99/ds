package list

import "github.com/HAo99/ds"

type List[T any] interface {
	ds.Container[T]

	Front() (T, error)
	Back() (T, error)

	Get(index int) (T, error)
	Set(index int, x T) error

	Insert(index int, x T) error
	PushFront(x T)
	PushBack(x T)

	Delete(index int) (T, error)
	PopFront() (T, error)
	PopBack() (T, error)
}

type ListX[T any] interface {
	FrontX() T
	BackX() T

	GetX(index int) T
	SetX(index int, x T)

	InsertX(index int, x T)
	DeleteX(index int) T
	PopFrontX() T
	PopBackX() T
}
