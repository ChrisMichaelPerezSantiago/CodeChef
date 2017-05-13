package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"
)

/*
   Problem: Turbo Sort
   Problem Code: TSORT

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
	start := time.Now()

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

	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
