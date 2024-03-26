package main

import (
	"errors"
	"fmt"
)

type Stack[T any] struct {
	item *item[T]
}

type item[T any] struct {
	Data T
	next *item[T]
}

type Stacker[T any] interface {
	Push(any)
	Pop() (T, error)
	IsEmpty() bool
}

func NewStack[T any]() Stacker[T] {
	// Если использовать такой конструктор, то
	// у нас стек никогда не будет nil

	return &Stack[T]{}
}

func (s *Stack[T]) Push(data any) {
	//Тут у параметра тип any для корректного сравнения с nil

	if s == nil || data == nil {
		return
	}
	s.item = &item[T]{data.(T), s.item}
}

func (s *Stack[T]) Pop() (data T, err error) {
	if s.IsEmpty() {
		return data, errors.New("stack is empty")
	}
	data = s.item.Data
	s.item = s.item.next
	return data, nil
}

func (s *Stack[T]) IsEmpty() bool {
	if s == nil || s.item == nil {
		return true
	} else {
		return false
	}
}
func main() {
	st := NewStack[any]()
	st.Push(func() { fmt.Println("We can do it") })
	st.Push(1.1)
	st.Push(nil)
	st.Push([...]int{1, 2, 3})
	st.Push("Goooood morning Vietnam")
	for !st.IsEmpty() {
		if val, err := st.Pop(); err == nil {
			fmt.Printf("Item: %v\n", val)
		}
	}
	_, err := st.Pop()
	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
	}
}
