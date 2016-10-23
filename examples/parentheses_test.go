package examples

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestIsBalanced(t *testing.T){
	balanced := "[()]{}{[()()]()}"
	balanced2 := "((((())))){[[]][]}"
	unbalanced := "[(])"
	unbalanced2 := "[{("
	empty := ""
	assert.Equal(t, true,IsBalanced(balanced))
	assert.Equal(t, true,IsBalanced(balanced2))
	assert.Equal(t, false,IsBalanced(unbalanced))
	assert.Equal(t, false,IsBalanced(unbalanced2))
	assert.Equal(t, true,IsBalanced(empty))
}