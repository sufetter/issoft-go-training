package main

import (
	"fmt"
	"time"
)

//В данном файле приведено несколько вариантов решения,
//а именно 6шт. Я бы вынес победителя в отдельный файл, но покеты/модули пока недоступны.
//Смотрите на whereAsync4 как на главный вариант,
//по результатам тестирования (в большинстве случаев) он оказался самым быстрым.
//Стоит отметить, что результаты очень сильно зависят от функции проверки, мне было очень тяжело
//определить какой вариант оставить, ведь при проверке i%2==0 и i%7==1 результаты отличаются очень сильно.
//Я очень люблю оптимизировать скорость работы программ, но тут я не смог конкретно определить,
//как это делать, ввиду того, что результаты даже при неизменности функции проверки могут сильно отличаться.

func whereBasic(nums []int, check func(int) bool) []int {
	length := len(nums)
	if length == 0 || check == nil {
		return nil
	}
	retNums := make([]int, 0, length)
	for i := 0; i < length; i++ {
		if check(nums[i]) {
			retNums = append(retNums, nums[i])
		}
	}

	return retNums
}

func whereAsync[T any](slice []T, checker func(T) bool) (sorted []T) {
	length := len(slice)

	if length == 0 || checker == nil {
		return
	}

	ch := make(chan []T, 8)

	for i := 0; i < 8; i++ {
		go func(elements []T, check func(T) bool, ch chan<- []T) {
			var filtered []T
			for _, elem := range elements {
				if check(elem) {
					filtered = append(filtered, elem)
				}
			}
			ch <- filtered
		}(slice[i*length/8:(i+1)*length/8], checker, ch)
	}

	for i := 0; i < 8; i++ {
		filtered := <-ch
		sorted = append(sorted, filtered...)
	}

	return
}

func whereAsyncV2[T any](slice []T, checker func(T) bool) []T {
	length := len(slice)

	if length == 0 || checker == nil {
		return nil
	}

	ch := make(chan []T, 8)
	for i := 0; i < 8; i++ {
		go func(elements []T, check func(T) bool, ch chan<- []T) {
			var filtered []T
			for j := 0; j < len(elements); j++ {
				if check(elements[j]) {
					filtered = append(filtered, elements[j])
				}
			}
			ch <- filtered
		}(slice[i*length/8:min((i+1)*length/8, length)], checker, ch)
	}

	sorted := make([]T, 0, 8)
	for i := 0; i < 8; i++ {
		filtered := <-ch
		sorted = append(sorted, filtered...)
	}

	return sorted
}

func whereAsyncV3[T any](slice []T, checker func(T) bool) []T {
	length := len(slice)

	if length == 0 || checker == nil {
		return nil
	}

	ch := make(chan []T, 8)
	for i := 0; i < 8; i++ {
		go func(elements []T, check func(T) bool, ch chan<- []T) {
			var filtered []T
			for j := 0; j < len(elements); j++ {
				if check(elements[j]) {
					filtered = append(filtered, elements[j])
				}
			}
			ch <- filtered
		}(slice[i*length/8:min((i+1)*length/8, length)], checker, ch)
	}

	sorted := make([]T, length)
	for i := 0; i < 8; i++ {
		filtered := <-ch
		for j, v := range filtered {
			sorted[i*length/8+j] = v
		}
	}

	return sorted
}

func whereAsyncV4[T any](slice []T, checker func(T) bool) []T {
	length := len(slice)

	if length == 0 || checker == nil {
		return nil
	}

	ch := make(chan []T, 8)
	done := make(chan bool, 8)
	for i := 0; i < 8; i++ {
		go func(elements []T, check func(T) bool, ch chan<- []T, done chan<- bool) {
			filtered := make([]T, 0, len(elements))
			for _, element := range elements {
				if check(element) {
					filtered = append(filtered, element)
				}
			}
			ch <- filtered
			done <- true
		}(slice[i*length/8:min((i+1)*length/8, length)], checker, ch, done)
	}

	go func() {
		for i := 0; i < 8; i++ {
			<-done
		}
		close(ch)
	}()

	sorted := make([]T, 0, length)
	for filtered := range ch {
		sorted = append(sorted, filtered...)
	}

	return sorted
}

func whereAsyncV5[T any](slice []T, checker func(T) bool) []T {
	length := len(slice)

	if length == 0 || checker == nil {
		return nil
	}

	ch := make(chan []T, 8)
	done := make(chan bool, 8)
	for i := 0; i < 8; i++ {
		go func(elements []T, check func(T) bool, ch chan<- []T, done chan<- bool) {
			count := 0
			for j := 0; j < len(elements); j++ {
				if check(elements[j]) {
					count++
				}
			}
			filtered := make([]T, count)
			count = 0
			for j := 0; j < len(elements); j++ {
				if check(elements[j]) {
					filtered[count] = elements[j]
					count++
				}
			}
			ch <- filtered
			done <- true
		}(slice[i*length/8:min((i+1)*length/8, length)], checker, ch, done)
	}

	go func() {
		for i := 0; i < 8; i++ {
			<-done
		}
		close(ch)
	}()

	sorted := make([]T, 0, length)
	for filtered := range ch {
		sorted = append(sorted, filtered...)
	}

	return sorted
}

func whereAsyncV6[T any](slice []T, checker func(T) bool) []T {
	length := len(slice)

	if length == 0 || checker == nil {
		return nil
	}

	ch := make(chan []T, 8)
	done := make(chan bool, 8)
	for i := 0; i < 8; i++ {
		go func(elements []T, check func(T) bool, ch chan<- []T, done chan<- bool) {
			count := 0
			for _, element := range elements {
				if check(element) {
					count++
				}
			}
			filtered := make([]T, count)
			count = 0
			for _, element := range elements {
				if check(element) {
					filtered[count] = element
					count++
				}
			}
			ch <- filtered
			done <- true
		}(slice[i*length/8:min((i+1)*length/8, length)], checker, ch, done)
	}

	go func() {
		for i := 0; i < 8; i++ {
			<-done
		}
		close(ch)
	}()

	sorted := make([]T, 0, length)
	for filtered := range ch {
		sorted = append(sorted, filtered...)
	}

	return sorted
}

func measureTime[T any](f func([]T, func(T) bool) []T, slice []T, checker func(T) bool) {
	start := time.Now()
	f(slice, checker)
	fmt.Printf("Function took %v\n", time.Since(start))
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	if n%2 == 0 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(277 % 4)
	nums := make([]int, 100000000)
	for i := 0; i < 100000000; i++ {
		nums[i] = i
	}
	check := func(i int) bool { return i%4*2*3/5 == 1 }

	functions := []func([]int, func(int) bool) []int{
		whereAsync[int],
		whereAsyncV2[int],
		whereAsyncV3[int],
		whereAsyncV4[int],
		whereAsyncV5[int],
		whereAsyncV6[int],
	}

	for i, function := range functions {
		fmt.Print(i + 1)
		measureTime(function, nums, check)
	}
	fmt.Printf("\nOne thread ")
	measureTime(whereBasic, nums, check)

	strings := make([]string, 40000000)
	for i := 0; i < 40000000; i++ {
		strings[i] = fmt.Sprintf("string%d", i)
	}
	checkString := func(s string) bool {
		for _, runeValue := range s {
			if isPrime(int(runeValue)) {
				return true
			}
		}
		return false
	}
	fmt.Println("\nWith strings")
	functionsString := []func([]string, func(string) bool) []string{
		whereAsync[string],
		whereAsyncV2[string],
		whereAsyncV3[string],
		whereAsyncV4[string],
		whereAsyncV5[string],
		whereAsyncV6[string],
	}
	for i, function := range functionsString {
		fmt.Print(i + 1)
		measureTime(function, strings, checkString)
	}
}
