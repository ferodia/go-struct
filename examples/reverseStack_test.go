package examples

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/ferodia/go-struct"
)


func TestReverse(t *testing.T){
	stack := gostruct.CreateStack(1, 2, 3, 4, 5)
	assert.Equal(t,stack.ToString(), "5 4 3 2 1 ")
	reversedStack := reverse(&stack)
	assert.Equal(t,reversedStack.ToString(), "1 2 3 4 5 ")
}