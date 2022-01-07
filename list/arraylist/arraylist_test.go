package arraylist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayList(t *testing.T) {
	l := New[int]()
	assert.Equal(t, 0, l.Len())
	assert.True(t, l.Empty())

	var err error
	_, err = l.Front()
	assert.ErrorIs(t, ErrIndexOutOfRange, err)
	_, err = l.Back()
	assert.ErrorIs(t, ErrIndexOutOfRange, err)
	_, err = l.Get(0)
	assert.ErrorIs(t, ErrIndexOutOfRange, err)
	err = l.Set(0, -1)
	assert.ErrorIs(t, ErrIndexOutOfRange, err)

	assert.PanicsWithError(t, ErrIndexOutOfRange.Error(), func() {
		l.GetX(0)
	})

	assert.PanicsWithError(t, ErrIndexOutOfRange.Error(), func() {
		l.SetX(0, -1)
	})

	assert.PanicsWithError(t, ErrIndexOutOfRange.Error(), func() {
		l.FrontX()
	})

	assert.PanicsWithError(t, ErrIndexOutOfRange.Error(), func() {
		l.BackX()
	})

	assert.PanicsWithError(t, ErrIndexOutOfRange.Error(), func() {
		l.DeleteX(0)
	})

	assert.PanicsWithError(t, ErrIndexOutOfRange.Error(), func() {
		l.PopBackX()
	})

	assert.PanicsWithError(t, ErrIndexOutOfRange.Error(), func() {
		l.PopFrontX()
	})

	l.PushFront(1)
	l.PushFront(2)
	l.PushFront(3)
	l.PushBack(4)
	l.PushBack(5)
	l.PushBack(6)
	assert.False(t, l.Empty())
	assert.Equal(t, 6, l.Len())
	assert.Equal(t, 3, l.GetX(0))
	assert.Equal(t, 2, l.GetX(1))
	assert.Equal(t, 1, l.GetX(2))
	assert.Equal(t, 4, l.GetX(3))
	assert.Equal(t, 5, l.GetX(4))
	assert.Equal(t, 6, l.GetX(5))
	assert.Equal(t, 3, l.FrontX())
	assert.Equal(t, 6, l.BackX())

	assert.PanicsWithError(t, ErrIndexOutOfRange.Error(), func() {
		l.InsertX(7, 10)
	})

	l.SetX(2, 5)
	assert.Equal(t, 5, l.GetX(2))

	assert.Equal(t, 4, l.DeleteX(3))
	assert.Equal(t, 3, l.PopFrontX())
	assert.Equal(t, 4, l.Len())
	assert.Equal(t, 6, l.PopBackX())
	assert.Equal(t, 3, l.Len())

}
