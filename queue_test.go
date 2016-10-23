package gostruct

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestQueueEmpty(t *testing.T){
	q := Queue{head:nil, length:0}
	assert.Equal(t, true, q.isEmpty())
	assert.Equal(t, 0, q.Size())
}

func TestEnqueue(t *testing.T){
	q := Queue{head:nil, length:0}
	q.enqueue(5)
	assert.Equal(t, false, q.isEmpty())
	assert.Equal(t, "5 ", q.toString())
}

func TestDequeue(t *testing.T){
	q := Queue{head:nil, length:0}
	q.enqueue("World")
	q.enqueue("Hello")
	assert.Equal(t, "World Hello ", q.toString())
	assert.Equal(t, "World", q.dequeue())
	assert.Equal(t, "Hello ", q.toString())
	assert.Equal(t, "Hello", q.dequeue())
	assert.Equal(t, nil, q.dequeue())
}


func BenchmarkQueue(b *testing.B) {
	// run the Enqueue function b.N times
	q := Queue{head:nil, length:0}
	for n := 0; n < b.N; n++ {
		q.enqueue(n)
		q.dequeue()
	}
}





