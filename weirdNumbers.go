package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	start := time.Now()
	var numbers []int

	for i := 1; i < 1001; i++ {
		if tapOrder(i) {
			numbers = append(numbers, i)
		}
	}
	for _, v := range numbers {
		if !sumCheck(v, findDivisors(v)) {
			println("I will be ordering the beer with the tap number:", v)
		}
	}
	elapsed := time.Since(start)
	fmt.Println("Execution time:", elapsed)
}

func perfectSquare(a int) bool {
	root := int(math.Sqrt(float64(a)))
	return (root * root) == a
}

// Checks if any combination of divisors' sum is equal to the number (condition 2)
func sumCheck(number int, divisors []int) bool {
	n := len(divisors)
	for i := 0; i < (1 << uint(n)); i++ {
		var combination []int
		for y := 0; y < n; y++ {
			// is the bit "on" in this number?
			if i&(1 << uint(y)) != 0 {
				//then add it to the combination
				combination = append(combination, divisors[y])
			}
		}
		if sumOfDivisors(combination) == number {
			return true
		}
	}
	return false
}

//
func tapOrder(number int) bool {
	if perfectSquare(number) {
		return false
	}
	divisors := findDivisors(number)
	if len(divisors) < 3 {
		// prime number
		return false
	}
	if sumOfDivisors(divisors) < number {
		return false
	}
	//if sumCheck(number, divisors){
	//	return false
	//}
	for i := len(divisors)-1; i > -1; i-- {
		divisors = findDivisors(number)
		if eliminateNumbers(divisors, i, number) {
			return false
		}

	}
	return true
}

// eliminateNumbers and eliminateNumbersHelper are 2 functions to recursively check if sum of divisors is equal to the numbers
// I added these functions to speed up the process, these functions do not try every combination like sumCheck function, but a lot of
// numbers are eliminated through these functions.
func eliminateNumbers(divisors []int, index int, number int) bool {
	number = number - divisors[index]
	return eliminateNumbersHelper(append(divisors[:index], divisors[index+1:]...), number)
}

func eliminateNumbersHelper(divisors []int, number int) bool {
	if number == 0 {
		return true
	}
	if number < 0 || len(divisors) == 0 {
		return false
	}
	for i := len(divisors)-1; i > -1; i-- {
		number = number - divisors[i]
		if eliminateNumbersHelper(append(divisors[:i], divisors[i+1:]...), number) {
			return true
		}
		return false
	}
	return false
}
// sum of number's divisors
func sumOfDivisors(divisors []int) int {
	sum := 0
	for _, v := range divisors {
		sum += v
	}
	return sum
}

// O(n/2), finding the divisors
func findDivisors(number int) []int {

	counter := int(number/2) + 1
	divisors := make([]int, 0)
	for i := 1; i < counter; i++ {
		if number%i == 0 {
			divisors = append(divisors, i)
		}
	}
	return divisors
}
