package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPushFront(t *testing.T) {
	assert := assert.New(t)
	ll := New()

	ll.PushFront("test1")
	assert.Equal(1, ll.Size())
	el, err := ll.Get(0)
	assert.Nil(err)
	assert.Equal("test1", el.Value)

	ll.PushFront("test2")
	assert.Equal(2, ll.Size())
	el, err = ll.Get(0)
	assert.Nil(err)
	assert.Equal("test2", el.Value)
}

func TestPopFront(t *testing.T) {
	assert := assert.New(t)
	ll := New()

	ll.PushFront("test1")
	ll.PushFront("test2")

	el := ll.PopFront()
	assert.Equal(1, ll.Size())
	assert.Equal("test2", el.Value)

	el = ll.PopFront()
	assert.Equal(0, ll.Size())
	assert.Equal("test1", el.Value)
}

func TestFind(t *testing.T) {
	assert := assert.New(t)
	ll := New()

	el := ll.Find("test")
	assert.Nil(el)

	ll.PushFront("test1")
	ll.PushFront("test2")

	el = ll.Find("test")
	assert.Nil(el)

	el = ll.Find("test2")
	assert.Equal("test2", el.Value)
}

func TestRemove(t *testing.T) {
	assert := assert.New(t)
	ll := New()

	err := ll.Remove(0)
	assert.EqualError(err, ErrInvalidIndex.Error())

	ll.PushFront("test")
	err = ll.Remove(0)
	assert.Nil(err)
	assert.Equal(0, ll.Size())

	ll.PushFront("test1")
	ll.PushFront("test2")
	err = ll.Remove(0)
	assert.Nil(err)
	assert.Equal(1, ll.Size())

	ll.PushFront("test3")
	ll.PushFront("test4")
	err = ll.Remove(2)
	assert.Nil(err)
	assert.Equal(2, ll.Size())
}

func TestGet(t *testing.T) {
	assert := assert.New(t)
	ll := New()

	el, err := ll.Get(0)
	assert.Nil(el)
	assert.EqualError(err, ErrInvalidIndex.Error())

	ll.PushFront("test1")
	ll.PushFront("test2")

	el, err = ll.Get(1)
	assert.Nil(err)
	assert.Equal("test1", el.Value)
}

func TestIndexOf(t *testing.T) {
	assert := assert.New(t)
	ll := New()

	index := ll.IndexOf("test")
	assert.Equal(-1, index)

	ll.PushFront("test1")
	ll.PushFront("test2")

	index = ll.IndexOf("test")
	assert.Equal(-1, index)

	index = ll.IndexOf("test1")
	assert.Equal(1, index)
}

func TestContaints(t *testing.T) {
	assert := assert.New(t)
	ll := New()

	containts := ll.Contains("test")
	assert.False(containts)

	ll.PushFront("test1")
	ll.PushFront("test2")

	containts = ll.Contains("test")
	assert.False(containts)

	containts = ll.Contains("test1")
	assert.True(containts)
}
