package main

import "fmt"

type num []int

func main() {
	nums := num{}

	for i := 1; i <= 10; i++ {
		nums = append(nums, i)
	}

	checkEvenOdd(nums)
}

func checkEvenOdd(nums num) {
	for _, n := range nums {
		if n%2 == 0 {
			fmt.Println(n, "is even")
		} else {
			fmt.Println(n, "is odd")
		}
	}
}
