package main

import "fmt"

func CountDown(n int) {
	if n <= 0 {
		return
	}
	fmt.Println(n)
	CountDown(n - 1)
}

func CountUp(n int) {
	if n <= 0 {
		return
	}
	CountUp(n - 1)
	fmt.Println(n)
}

func Subtract(n int) int {
	if n <= 0 {
		return 0
	}
	//	return n - Subtract(n-1)
	return Subtract(n-1) - n
}

func Power5(n int) float64 {
	if n == 0 {
		return 1
	}

	if n < 0 {
		return 1 / Power5(-n)
	}

	return 5 * Power5(n-1)
}

func main() {
	CountDown(3)
	CountUp(3)
}
