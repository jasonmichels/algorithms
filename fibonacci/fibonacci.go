package main

import (
	"fmt"
	"runtime"
	"time"
)

func fibonacciSlow(n int) int {
	if n <= 1 {
		return n
	}

	return fibonacciSlow(n - 1) + fibonacciSlow(n - 2)
}

func fibonacciFast(n int) int {
	if n <= 1 {
		return n
	}

	numbers := make([]int, n + 1)
	numbers[0] = 0
	numbers[1] = 1

	for i := 2; i <= n; i++ {
		numbers[i] = numbers[i - 1] + numbers[i - 2]
	}

	return numbers[n]
}

func printSlow() {
	fmt.Printf("Slow: %d \n", fibonacciSlow(15))
}

func printFast() {
	fmt.Printf("Fast: %d \n", fibonacciFast(15))
}

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	go printSlow()
	go printFast()

	time.Sleep(2 * time.Minute)
}