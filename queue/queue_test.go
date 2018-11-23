package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnqueue(t *testing.T) {
	assert := assert.New(t)
	q := New()
	execepted := []interface{}{"test1", "test2"}

	q.Enqueue("test1", "test2")

	assert.Len(q.elements, 2)
	assert.Equal(execepted, q.elements)
	assert.Equal(execepted[0], q.elements[0])
}
func TestDequeue(t *testing.T) {
	assert := assert.New(t)
	q := New()

	q.Enqueue("test1", "test2")

	firstElement, err := q.Dequeue()
	assert.Nil(err)
	assert.Len(q.elements, 1)
	assert.Equal("test1", firstElement)
	assert.Equal([]interface{}{"test2"}, q.elements)

	secondElement, err := q.Dequeue()
	assert.Nil(err)
	assert.Len(q.elements, 0)
	assert.Equal("test2", secondElement)
	assert.Equal([]interface{}{}, q.elements)

	_, err = q.Dequeue()
	assert.EqualError(err, ErrEmptyQueue.Error())
}

func TestPeek(t *testing.T) {
	assert := assert.New(t)
	q := New()
	testElements := []interface{}{"test1", "test2"}

	_, err := q.Peek()
	assert.EqualError(err, ErrEmptyQueue.Error())

	q.Enqueue(testElements[0])
	element, err := q.Peek()
	assert.Nil(err)
	assert.Len(q.elements, 1)
	assert.Equal(testElements[0], element)
	assert.Equal(testElements[:1], q.elements)

	q.Enqueue(testElements[1])
	element, err = q.Peek()
	assert.Nil(err)
	assert.Len(q.elements, 2)
	assert.Equal(testElements[0], element)
	assert.Equal(testElements, q.elements)
}

func TestSize(t *testing.T) {
	assert := assert.New(t)
	q := New()

	size := q.Size()
	assert.Zero(size)

	q.Enqueue("test1", "test2")
	size = q.Size()
	assert.Equal(2, size)

	q.Dequeue()
	size = q.Size()
	assert.Equal(1, size)
}

func TestIsEmpty(t *testing.T) {
	assert := assert.New(t)
	q := New()

	isEmpty := q.IsEmpty()
	assert.True(isEmpty)

	q.Enqueue("test")
	isEmpty = q.IsEmpty()
	assert.False(isEmpty)
}
