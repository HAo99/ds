package queue

import "github.com/HAo99/ds"

type Queue[T any] interface {
	ds.Container[T]

	Peek() (T, error)
	Enqueue(x T)
	Dequeue() (T, error)
}

type QueueX[T any] interface {
	Peek() (T, error)
	EnqueueX(x T)
	DequeueX() (T, error)
}
