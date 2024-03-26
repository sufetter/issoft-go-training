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
	Push(T)
	Pop() (T, error)
	IsEmpty() bool
}

func NewStack[T any]() Stacker[T] {
	// Если использовать такой конструктор, то
	// у нас стек никогда не будет nil

	return &Stack[T]{}
}

func (s *Stack[T]) Push(data T) {
	if s == nil {
		return
	}
	s.item = &item[T]{data, s.item}
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
	// На данный момент параметр data имеет тип T, что не всегда может быть nil
	// и тогда произойдет паника при сравнении с nil (если тип T не может быть nil).
	// Можно тут data (тот что параметр в методах Push и Pop) сделать типом any,
	// а не T и тогда можно спокойно обработать nil, но я не знаю хорошо так делать или нет.

	st := NewStack[any]()
	st.Push(func() { fmt.Println("We can do it") })
	st.Push(1.1)
	//st.Push(nil) плохой сценарий
	st.Push([...]int{1, 2, 3})
	st.Push("Goooood morning Vietnam")
	for !st.IsEmpty() {
		if val, err := st.Pop(); err == nil {
			fmt.Printf("Item: %v\n", val)
		}
	}
	_, err := st.Pop()
	if err != nil {
		fmt.Println(err.Error())
	}
}
