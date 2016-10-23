package gostruct

import (
	"strconv"
)

type Stack struct {
	head * item
	length int
}

type item struct {
	value interface{}
	next *item
}

func NewStack()(Stack){
	return Stack{head:nil, length:0}
}

func CreateStack(values...interface{})(Stack){
	stack := Stack{head:nil, length:0}
	for _,e := range values{
		stack.Push(e)
	}
	return stack
}

func (s * Stack)Push(v interface{}){
	itm := item{v, nil}
	if s.IsEmpty(){
		s.head = &itm
	}else{
		tmp := s.head
		s.head = &itm
		s.head.next = tmp
	}
	s.length +=1
}

func (s * Stack)Pop() (interface{}){
	if !s.IsEmpty(){
		curr := s.head
		s.head = s.head.next
		s.length -= 1
		return curr.value
	}
	return nil
}

func (s Stack) Top() (interface{}){
	return s.Pop()
}

func (s Stack)IsEmpty()bool{
	return s.length == 0
}

func (s Stack) Size() int {
	return s.length
}

func (s * Stack)ToString() (values string ){
	curr := s.head
	for curr != nil {
		_, ok := curr.value.(string)
		if ok == true{
			values += curr.value.(string) +" "
		}else{
			values += strconv.Itoa(curr.value.(int)) +" "
		}
		curr = curr.next
	}
	return
}



