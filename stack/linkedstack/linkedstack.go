package linkedstack

import (
	"errors"

	"github.com/HAo99/ds/list/linkedlist"
	"github.com/HAo99/ds/stack"
)

var (
	_ stack.Stack[struct{}]  = (*LinkedStack[struct{}])(nil)
	_ stack.StackX[struct{}] = (*LinkedStack[struct{}])(nil)
)

var (
	ErrEmptyStack = errors.New("empty stack")
)

type LinkedStack[T any] struct {
	list linkedlist.LinkedList[T]
}

func New[T any]() *LinkedStack[T] {
	return &LinkedStack[T]{
		list: *linkedlist.New[T](),
	}
}

func (s *LinkedStack[T]) Len() int {
	return s.list.Len()
}

func (s *LinkedStack[T]) Empty() bool {
	return s.list.Empty()
}

func (s *LinkedStack[T]) Push(x T) {
	s.list.PushBack(x)
}

func (s *LinkedStack[T]) Peek() (T, error) {
	r, err := s.list.Back()
	if err != nil {
		var zero T
		return zero, ErrEmptyStack
	}
	return r, nil
}

func (s *LinkedStack[T]) PeekX() T {
	r, err := s.list.Back()
	if err != nil {
		panic(err)
	}
	return r
}

func (s *LinkedStack[T]) Pop() (T, error) {
	r, err := s.list.PopBack()
	if err != nil {
		var zero T
		return zero, ErrEmptyStack
	}
	return r, nil
}

func (s *LinkedStack[T]) PopX() T {
	r, err := s.list.PopBack()
	if err != nil {
		panic(err)
	}
	return r
}
