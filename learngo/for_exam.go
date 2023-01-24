package main

import (
	"fmt"
)

func superAdd(numbers ...int) int {
	total := 0
	// for _, number := range numbers {
	// 	total += number
	// }
	for i := 0; i < len(numbers); i++ {
		total += numbers[i]
	}
	return total
}

func canIDrink(age int) bool {
	defer fmt.Println("나이판독")
	// if koreanAge := age + 2; koreanAge < 18 {
	// 	return false
	// }
	// return true
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	// var total int
	// total = superAdd(1, 2, 3, 4, 5, 6)
	// fmt.Println(total)
	// fmt.Println(canIDrink(16))
	foods := []string{"kimchi", "ramen"}
	solda := person{name: "sol", age: 28, favFood: foods}
	fmt.Println(solda)
}
