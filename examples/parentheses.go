package examples

import (
	"github.com/ferodia/go-struct"
)

func isOpen(input string) (bool){
	return( input == "(" ||
			input == "[" ||
			input == "{")
}


func match(e1 string, e2 string) (bool){
	return( (e1 == "{" && e2 =="}") ||
		(e1 == "(" && e2 ==")") ||
		(e1 == "[" && e2 =="]") )

}

func IsBalanced(input string) bool {
	s :=  gostruct.NewStack()
	for _, elem := range input {
		if isOpen(string(elem)){
			s.Push(string(elem))
		} else{
			if s.IsEmpty(){
				return false
			}
			elem2 := s.Pop()
			if !match(elem2.(string), string(elem)){
				return false
			}
		}
	}
	if s.IsEmpty(){
		return true
	}
	return false
}


