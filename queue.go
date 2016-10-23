package gostruct

import (
	"strconv"
)

type Queue struct {
	head *item
	tail *item
	length int
}


func (q *Queue)enqueue(value interface{}) {
	itm := item{value, nil}
	tmp := q.tail
	q.tail = &itm
	if q.isEmpty(){
		q.head = q.tail
	} else{
		tmp.next = q.tail
	}
	q.length += 1
}

func (q * Queue)dequeue() (val interface{}) {
	if !q.isEmpty() {
		val = q.head.value
		q.head = q.head.next
		q.length -= 1
		return
	}
	if q.isEmpty(){
		q.tail = nil
	}

	return val
}

func (q Queue)isEmpty()bool{
	return q.length == 0
}

func (q Queue) Size() int {
	return q.length
}


func (q * Queue)toString() (values string ){
	curr := q.head
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

