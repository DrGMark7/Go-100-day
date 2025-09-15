package main

import (
	"fmt"
)

func main() {
	var a []int
	sum := 0

	for i := 0; i < 5; i++ {
		var n int
		fmt.Scan(&n)
		a = append(a, n)
		fmt.Println(a)
		sum += n
	}	

	fmt.Printf("Total Sum is %d", sum)
}