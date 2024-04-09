package main

import (
	"fmt"
	"time"
)

//Я очень люблю оптимизировать скорость работы программ, но тут я не смог конкретно определить,
//как это делать, ввиду того, что результаты даже при неизменности функции проверки могут сильно отличаться.
//Так, если использовать 2 разные функции проверки, то самая медленная функция whereAsync умудрялась стать
//самой быстрой, а самая быстрая наоборот.
//У меня было 6 вариантов реализации данной функции, но чем дольше я их сравнивал, тем больше терялся в цифрах
//производительности.
//В результате я оставил ту, что в среднем (но не всегда и не со всеми типами данных) показала лучшие результаты.
//Также уточню, что создавал срезы больших размеров в функциях для оптимизации именно скорости работы, а не
//потребления памяти, если создать маленькие или 0 срезы, то количество аллокаций может быть огромным.

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

func whereAsync[T any](slice []T, checker func(T) bool) []T {
	length := len(slice)

	if length == 0 || checker == nil {
		return nil
	}

	ch := make(chan []T, 8)
	for i := 0; i < 8; i++ {
		go func(elements []T, check func(T) bool, ch chan<- []T) {
			filtered := make([]T, 0, len(elements))
			for _, element := range elements {
				if check(element) {
					filtered = append(filtered, element)
				}
			}
			ch <- filtered
		}(slice[i*length/8:min((i+1)*length/8, length)], checker, ch)
	}

	sorted := make([]T, 0, length)
	for i := 0; i < 8; i++ {
		sorted = append(sorted, <-ch...)
	}

	return sorted
}

func measureTime[T any](f func([]T, func(T) bool) []T, slice []T, checker func(T) bool) {
	if f == nil || checker == nil || slice == nil {
		fmt.Println("Check parameters")
		return
	}
	start := time.Now()
	f(slice, checker)
	fmt.Printf("Function took %v\n", time.Since(start))
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
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
	//Тут все в циклах и срезах,
	//тк при тестировании удобно было запускать сразу все функции.
	//PC: Ryzen 5 3600, 16gb DDR4 3200; WSL - Ubuntu 22.04

	nums := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		nums[i] = i
	}

	check := func(i int) bool { return i%2 == 0 }

	functions := []func([]int, func(int) bool) []int{
		whereAsync[int],
	}

	//На моем компьютере в среднем 90µs
	fmt.Println("With small slice")
	for i, function := range functions {
		fmt.Print(i + 1)
		measureTime(function, nums, check)
	}
	fmt.Printf("One thread ")
	//24µs
	measureTime(whereBasic, nums, check)

	nums = make([]int, 100000000)
	for i := 0; i < 100000000; i++ {
		nums[i] = i
	}

	fmt.Println("\nNow with large slice")

	//120ms
	for i, function := range functions {
		fmt.Print(i + 1)
		measureTime(function, nums, check)
	}
	fmt.Printf("One thread ")

	//250ms
	measureTime(whereBasic, nums, check)

	strings := make([]string, 40000000)
	for i := 0; i < 40000000; i++ {
		strings[i] = fmt.Sprintf("%d", i)
	}
	checkString := func(s string) bool {
		for _, runeValue := range s {
			if !isPrime(int(runeValue)) {
				return false
			}
		}
		return true
	}
	fmt.Println("\nWith strings")
	functionsString := []func([]string, func(string) bool) []string{
		whereAsync[string],
	}

	//~100ms разброс очень сильный
	for i, function := range functionsString {
		fmt.Print(i + 1)
		measureTime(function, strings, checkString)
	}
}
