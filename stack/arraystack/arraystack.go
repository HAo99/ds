package arraystack

import (
	"errors"

	"github.com/HAo99/ds/list/arraylist"
	"github.com/HAo99/ds/stack"
)

var (
	_ stack.Stack[struct{}]  = (*ArrayStack[struct{}])(nil)
	_ stack.StackX[struct{}] = (*ArrayStack[struct{}])(nil)
)

var (
	ErrEmptyStack = errors.New("empty stack")
)

type ArrayStack[T any] struct {
	list arraylist.ArrayList[T]
}

func New[T any]() *ArrayStack[T] {
	return &ArrayStack[T]{
		list: *arraylist.New[T](),
	}
}

func FromArrayList[T any](list *arraylist.ArrayList[T]) *ArrayStack[T] {
	return &ArrayStack[T]{
		list: *list,
	}
}

func (s *ArrayStack[T]) Len() int {
	return s.list.Len()
}

func (s *ArrayStack[T]) Empty() bool {
	return s.list.Empty()
}

func (s *ArrayStack[T]) Push(x T) {
	s.list.PushBack(x)
}

func (s *ArrayStack[T]) Peek() (T, error) {
	r, err := s.list.Back()
	if err != nil {
		if errors.Is(err, arraylist.ErrIndexOutOfRange) {
			return r, ErrEmptyStack
		}
		panic(err)
	}
	return r, nil
}

func (s *ArrayStack[T]) PeekX() T {
	r, err := s.Peek()
	if err != nil {
		panic(err)
	}
	return r
}

func (s *ArrayStack[T]) Pop() (T, error) {
	r, err := s.list.PopBack()
	if err != nil {
		if errors.Is(err, arraylist.ErrIndexOutOfRange) {
			return r, ErrEmptyStack
		}
		panic(err)
	}
	return r, nil
}

func (s *ArrayStack[T]) PopX() T {
	r, err := s.Pop()
	if err != nil {
		panic(err)
	}
	return r
}
