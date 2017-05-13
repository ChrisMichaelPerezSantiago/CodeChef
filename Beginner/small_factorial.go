package main

import (
	"fmt"
	"math/big"
	"time"
)

/*
   Problem: Small Factorial
   Problem Code: FLOW018

   @Author:  Chris M. Perez
   @Date:    5/12/2017
*/

func factorial(args int64) *big.Int {
	if args < 0 {
		return big.NewInt(1)
	}
	if args == 0 {
		return big.NewInt(1)
	}

	bigInteger := big.NewInt(args)
	return bigInteger.Mul(bigInteger, factorial(args-1))
}

func main() {
	start := time.Now()

	var n int
	var args int64

	fmt.Scan(&n)
	for index := 0; index < n; index++ {
		fmt.Scan(&args)
		fmt.Println(factorial(args))
	}
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
