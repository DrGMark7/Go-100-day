package main

import "fmt"

// For Loop and While Loop in Go lang
func main() {

	for i := 0; i < 5; i++ {
		fmt.Println("For Loop:", i)
	}

	// While Loop in Go is implemented using a for loop without initialization and post statement
	N := 0
	for N != 0 {
		fmt.Println("While Loop:", N)
		N++
		if N == 5 {
			break
		}
	}

	list := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(list); i++ {
		fmt.Println("List item:", list[i])
	}

	for _, item := range list {
		fmt.Println("List item using range:", item)
	}

	for i, item := range list {
		fmt.Printf("List item with index %d: %d\n", i, item)
	}

}