package main

import "fmt"

/*
	Problem: Sum of Digits
	Problem Code: FLOW006

	@Author:  Chris M. Perez
	@Date:    5/12/2017
*/

func sumDigit(args int) (total int) {
	n := args

	for n != 0 {
		total += n % 10
		n /= 10
	}
	return total
}

func exe(n int) {
	var args int
	for index := 0; index < n; index++ {
		fmt.Scan(&args)
		fmt.Println(sumDigit(args))
	}
}

func main() {
	var n int

	fmt.Scan(&n)
	exe(n)
}
