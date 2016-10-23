package examples

import "github.com/ferodia/go-struct"


func reverse(s * gostruct.Stack) (gostruct.Stack){
	newStack := gostruct.NewStack()
	for !s.IsEmpty() {
		newStack.Push(s.Pop())
	}
	return newStack
}