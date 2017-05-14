package main

import (
	"fmt"
	"time"
)

/*
   Problem: Servant
   Problem Code: FLOW008

   @Author:  Chris M. Perez
   @Date:    5/13/2017
*/
func less(arg int) {
	if arg >= 10 {
		fmt.Println("-1")
	}

	if arg < 10 {
		fmt.Println("What an obedient servant you are!")
	}
}

func main() {
	start := time.Now()

	var n int
	var args int

	fmt.Scan(&n)
	for index := 0; index < n; index++ {
		fmt.Scan(&args)
		less(args)
	}

	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
