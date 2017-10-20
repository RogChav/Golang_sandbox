package main

import (
	"fmt"
)

func main() {
	numbers := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, current := range numbers {
		if current%2 != 0 {
			fmt.Println(current, " is odd")
		} else {
			fmt.Println(current, " is even")
		}
	}

}
