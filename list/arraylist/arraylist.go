package arraylist

import (
	"errors"

	"github.com/HAo99/ds/list"
)

var (
	_ list.List[struct{}]  = (*arrayList[struct{}])(nil)
	_ list.ListX[struct{}] = (*arrayList[struct{}])(nil)
)

var (
	ErrIndexOutOfRange = errors.New("index out of range")
)

type arrayList[T any] struct {
	data []T
	len  int
}

func New[T any]() *arrayList[T] {
	return &arrayList[T]{
		data: make([]T, 0),
		len:  0,
	}
}

func (l *arrayList[T]) Cap() int {
	return len(l.data)
}

func (l *arrayList[T]) Len() int {
	return l.len
}

func (l *arrayList[T]) Empty() bool {
	return l.len == 0
}

func (l *arrayList[T]) Front() (T, error) {
	if l.Empty() {
		var zero T
		return zero, ErrIndexOutOfRange
	}
	return l.data[0], nil
}

func (l *arrayList[T]) FrontX() T {
	r, err := l.Front()
	if err != nil {
		panic(err)
	}
	return r
}

func (l *arrayList[T]) Back() (T, error) {
	if l.Empty() {
		var zero T
		return zero, ErrIndexOutOfRange
	}
	return l.data[l.len-1], nil
}

func (l *arrayList[T]) BackX() T {
	r, err := l.Back()
	if err != nil {
		panic(err)
	}
	return r
}

func (l *arrayList[T]) Get(index int) (T, error) {
	if index >= l.len || index < 0 {
		var zero T
		return zero, ErrIndexOutOfRange
	}
	return l.data[index], nil
}

func (l *arrayList[T]) GetX(index int) T {
	r, err := l.Get(index)
	if err != nil {
		panic(err)
	}
	return r
}

func (l *arrayList[T]) Set(index int, x T) error {
	if index >= l.len || index < 0 {
		return ErrIndexOutOfRange
	}
	l.data[index] = x
	return nil
}

func (l *arrayList[T]) SetX(index int, x T) {
	err := l.Set(index, x)
	if err != nil {
		panic(err)
	}
}

func (l *arrayList[T]) Insert(index int, x T) error {
	if index > l.len || index < 0 {
		return ErrIndexOutOfRange
	}
	if l.len == l.Cap() {
		l.grow()
	}
	for i := l.len; i-1 >= index; i-- {
		l.data[i] = l.data[i-1]
	}
	l.data[index] = x
	l.len++
	return nil
}

func (l *arrayList[T]) InsertX(index int, x T) {
	err := l.Insert(index, x)
	if err != nil {
		panic(err)
	}
}

func (l *arrayList[T]) grow() {
	var newData []T
	if len(l.data) == 0 {
		newData = make([]T, 2)
	} else {
		newData = make([]T, len(l.data)<<1)
	}
	for i, v := range l.data {
		newData[i] = v
	}
	l.data = newData
}

func (l *arrayList[T]) PushBack(x T) {
	l.Insert(l.len, x)
}

func (l *arrayList[T]) PushFront(x T) {
	l.Insert(0, x)
}

func (l *arrayList[T]) Delete(index int) (T, error) {
	if index >= l.len || index < 0 {
		var zero T
		return zero, ErrIndexOutOfRange
	}
	ret := l.data[index]
	for i := index; i+1 < l.len; i++ {
		l.data[i] = l.data[i+1]
	}
	l.len--
	return ret, nil
}

func (l *arrayList[T]) DeleteX(index int) T {
	r, err := l.Delete(index)
	if err != nil {
		panic(err)
	}
	return r
}

func (l *arrayList[T]) PopFront() (T, error) {
	return l.Delete(0)
}

func (l *arrayList[T]) PopFrontX() T {
	r, err := l.PopFront()
	if err != nil {
		panic(err)
	}
	return r
}

func (l *arrayList[T]) PopBack() (T, error) {
	return l.Delete(l.len - 1)
}

func (l *arrayList[T]) PopBackX() T {
	r, err := l.PopBack()
	if err != nil {
		panic(err)
	}
	return r
}
