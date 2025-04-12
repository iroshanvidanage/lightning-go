package main

import "fmt"

func odd_even(i []int) {
	for _, n := range i {
		if n%2 == 0 {
			fmt.Println(n, "is Even.")
		} else {
			fmt.Println(n, "is Odd.")
		}
	}
}
