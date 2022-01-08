package linkedlist

import (
	"errors"

	"github.com/HAo99/ds/list"
)

var (
	_ list.List[struct{}]  = (*linkedList[struct{}])(nil)
	_ list.ListX[struct{}] = (*linkedList[struct{}])(nil)
)

var (
	ErrIndexOutOfRange = errors.New("index out of range")
)

type linkedList[T any] struct {
	root *node[T]
	len  int
}

type node[T any] struct {
	val  T
	next *node[T]
	prev *node[T]
}

func New[T any]() *linkedList[T] {
	root := &node[T]{}
	root.next = root
	root.prev = root
	return &linkedList[T]{
		root: root,
		len:  0,
	}
}

func (l *linkedList[T]) Len() int {
	return l.len
}

func (l *linkedList[T]) Empty() bool {
	return l.len == 0
}

func (l *linkedList[T]) Front() (T, error) {
	if l.Empty() {
		var zero T
		return zero, ErrIndexOutOfRange
	}
	return l.root.next.val, nil
}

func (l *linkedList[T]) FrontX() T {
	r, err := l.Front()
	if err != nil {
		panic(err)
	}
	return r
}

func (l *linkedList[T]) Back() (T, error) {
	if l.Empty() {
		var zero T
		return zero, ErrIndexOutOfRange
	}
	return l.root.prev.val, nil
}

func (l *linkedList[T]) BackX() T {
	r, err := l.Back()
	if err != nil {
		panic(err)
	}
	return r
}

func (l *linkedList[T]) Get(index int) (T, error) {
	if index >= l.len || index < 0 {
		var zero T
		return zero, ErrIndexOutOfRange
	}
	cur := l.root
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	return cur.next.val, nil
}

func (l *linkedList[T]) GetX(index int) T {
	r, err := l.Get(index)
	if err != nil {
		panic(err)
	}
	return r
}

func (l *linkedList[T]) Set(index int, x T) error {
	if index >= l.len || index < 0 {
		return ErrIndexOutOfRange
	}
	cur := l.root
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	cur.next.val = x
	return nil
}

func (l *linkedList[T]) SetX(index int, x T) {
	err := l.Set(index, x)
	if err != nil {
		panic(err)
	}
}

func (l *linkedList[T]) Insert(index int, x T) error {
	if index > l.len || index < 0 {
		return ErrIndexOutOfRange
	}

	cur := l.root
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	newNode := &node[T]{
		val:  x,
		next: cur.next,
		prev: cur,
	}
	newNode.next.prev = newNode
	cur.next = newNode
	l.len++
	return nil
}

func (l *linkedList[T]) InsertX(index int, x T) {
	err := l.Insert(index, x)
	if err != nil {
		panic(err)
	}
}

func (l *linkedList[T]) PushFront(x T) {
	newNode := &node[T]{
		val:  x,
		next: l.root.next,
		prev: l.root,
	}
	newNode.next.prev = newNode
	l.root.next = newNode
	l.len++
}

func (l *linkedList[T]) PushBack(x T) {
	newNode := &node[T]{
		val:  x,
		next: l.root,
		prev: l.root.prev,
	}
	newNode.prev.next = newNode
	l.root.prev = newNode
	l.len++
}

func (l *linkedList[T]) Delete(index int) (T, error) {
	if index >= l.len || index < 0 {
		var zero T
		return zero, ErrIndexOutOfRange
	}
	cur := l.root
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	ret := cur.next.val
	cur.next = cur.next.next
	cur.next.prev = cur
	l.len--
	return ret, nil
}

func (l *linkedList[T]) DeleteX(index int) T {
	r, err := l.Delete(index)
	if err != nil {
		panic(err)
	}
	return r
}

func (l *linkedList[T]) PopFront() (T, error) {
	if l.Empty() {
		var zero T
		return zero, ErrIndexOutOfRange
	}
	ret := l.root.next.val
	l.root.next = l.root.next.next
	l.root.next.prev = l.root
	l.len--
	return ret, nil
}

func (l *linkedList[T]) PopFrontX() T {
	r, err := l.PopFront()
	if err != nil {
		panic(err)
	}
	return r
}

func (l *linkedList[T]) PopBack() (T, error) {
	if l.Empty() {
		var zero T
		return zero, ErrIndexOutOfRange
	}
	ret := l.root.prev.val
	l.root.prev = l.root.prev.prev
	l.root.prev.next = l.root
	l.len--
	return ret, nil
}

func (l *linkedList[T]) PopBackX() T {
	r, err := l.PopBack()
	if err != nil {
		panic(err)
	}
	return r
}
