package main

import (
	"fmt"
	"slices"
)

func main()  {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(slices.Max(numbers))
}