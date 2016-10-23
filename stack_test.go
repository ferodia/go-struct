package gostruct

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestStackEmpty(t *testing.T){
	s := Stack{head:nil, length:0}
	assert.Equal(t, true, s.IsEmpty())


}

func TestPush(t *testing.T){
	s := Stack{head:nil, length:0}
	s.Push(5)
	assert.Equal(t, false, s.IsEmpty())
	assert.Equal(t, "5 ", s.ToString())
}

func TestPop(t *testing.T){
	s := Stack{head:nil, length:0}
	s.Push("World")
	s.Push("Hello")
	assert.Equal(t, "Hello World ", s.ToString())
	assert.Equal(t, "Hello", s.Pop())
	assert.Equal(t, "World ", s.ToString())
	assert.Equal(t, "World", s.Pop())
}

func TestTop(t * testing.T){
	s := Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	assert.Equal(t, 3, s.Top())
	assert.Equal(t, 3, s.Size())

}







