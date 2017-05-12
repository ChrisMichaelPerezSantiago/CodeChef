package main

import (
	"fmt"
	"strconv"
)

/*
	Problem: First and Last Digit
	Problem Code: FLOW004

	@Author:  Chris M. Perez
	@Date:    5/12/2017
*/

func sum(str string) (total int) {
	start, _ := strconv.Atoi(string([]rune(str)[0]))
	last, _ := strconv.Atoi(string([]rune(str)[len(str)-1]))

	total = start + last
	return total
}

func main() {
	var n int
	var str string

	fmt.Scan(&n)
	for index := 0; index < n; index++ {
		fmt.Scan(&str)
		fmt.Println(sum(str))
	}
}
