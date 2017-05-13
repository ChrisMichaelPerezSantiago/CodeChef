package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

/*
   Problem: Turbo Sort
   Problem Code: Problem Code: TSORT

   @Author:  Chris M. Perez
   @Date:    5/13/2017
*/

func sorting(args []int) {
	sort.Ints(args[:])

	for _, elements := range args {
		fmt.Println(elements)
	}
}

func main() {
	console := bufio.NewScanner(os.Stdin)
	var n int

	fmt.Scanf("%d\n", &n)
	input := make([]int, n)

	for index := 0; index < n; index++ {
		console.Scan()
		if n, err := strconv.Atoi(console.Text()); err == nil {
			input[index] = n
		}
	}
	sorting(input)
}
