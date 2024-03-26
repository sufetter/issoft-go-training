package main

import (
	"fmt"
	"reflect"
)

type Cloner interface {
	Clone() Cloner
}

type Human struct {
	iq byte
}

func (h *Human) Clone() Cloner {
	return &Human{iq: h.iq}
}

type CatGirl struct {
	state string
}

func (c *CatGirl) Clone() Cloner {
	return &CatGirl{state: c.state}
}

type Elf struct {
	age int
}

func (e *Elf) Clone() Cloner {
	return &Elf{e.age}
}

func sliceClone(slice []any) []any {
	newSlice := make([]interface{}, 0, len(slice))

	//Тут можно было бы просто if и все типы перечислить в 1 строчку, но я решил разбить на кейсы
	for _, item := range slice {
		val := reflect.ValueOf(item)
		switch val.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			newSlice = append(newSlice, val.Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			newSlice = append(newSlice, val.Uint())
		case reflect.Float32, reflect.Float64:
			newSlice = append(newSlice, val.Float())
		case reflect.Bool:
			newSlice = append(newSlice, val.Bool())
		case reflect.String:
			newSlice = append(newSlice, val.String())
		case reflect.Ptr:
			if cloner, ok := item.(Cloner); ok {
				newSlice = append(newSlice, cloner.Clone())
			}
		default:
			fmt.Printf("That type is not allowed: %v\n", val.Kind())
		}
	}
	return newSlice
}

func main() {
	slice := []any{
		nil,
		2003,
		uint(777),
		-42.42,
		true,
		"In 1989, nothing happened in Tienanmen Square in China.",
		&Human{iq: 123},
		&CatGirl{state: "I wanna be happy"},
		&Elf{age: 30},
	}
	clonedSlice := sliceClone(slice)
	fmt.Println(clonedSlice)
}
