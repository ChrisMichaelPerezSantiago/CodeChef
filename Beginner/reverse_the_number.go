package main

import (
   "fmt"
   "strconv"
)

/*
   Problem: Reverse The Number
   Problem Code: FLOW007

   @Author:  Chris M. Perez
   @Date:    5/12/2017
*/

func reversingNumbers(str string) int {
   var reverse string
   var rev int

   for _, elements := range str {
       reverse = string(elements) + reverse
       rev, _ = strconv.Atoi(reverse)
   }
   return rev
}

func main() {
   var n int
   var str string

   fmt.Scan(&n)
   for index := 0; index < n; index++ {
      fmt.Scan(&str)
      fmt.Println(reversingNumbers(str))
    }
}
