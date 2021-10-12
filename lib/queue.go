package lib

import "errors"

type IQueue interface {
	Push(int)
	Pop() (int, error)
}

type Queue struct {
	arr  []int
	size int
}

func (q *Queue) Push(a int) error {
	if len(q.arr) >= q.size {
		return errors.New("Queue overflow")
	}
	q.arr = append(q.arr, a)
	return nil
}

func (q *Queue) Pop() (int, error) {
	if len(q.arr) <= 0 {
		return -1, errors.New("Queue underflow")
	}
	var item int = q.arr[0]
	q.arr = q.arr[1:]
	return item, nil
}

func NewQueue(size int) *Queue {
	return &Queue{make([]int, 0), size}
}
