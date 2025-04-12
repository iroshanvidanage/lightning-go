package main

import "fmt"

func slice_int() {

	int_slice := []int{}
	for i := range 11 {
		int_slice = append(int_slice, i)
	}

	odd_even(int_slice)

}

func odd_even(i []int) {
	for _, n := range i {
		if n%2 == 0 {
			fmt.Println(n, "is Even.")
		} else {
			fmt.Println(n, "is Odd.")
		}
	}
}
