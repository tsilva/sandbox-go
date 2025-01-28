package main

import "fmt"

func main() {

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// BUG: this doesnt work
	for i := range make([]int, 3) {
		fmt.Println("range", i)
	}

	for {
		fmt.Println("loop")
		break
	}

	// BUG: this doesnt work
	for n := range make([]int, 6) {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
