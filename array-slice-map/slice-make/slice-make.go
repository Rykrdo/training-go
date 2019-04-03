package main

import "fmt"

func main() {
	s := make([]int, 10)
	s[9] = 21

	fmt.Println(cap(s))

	s = make([]int, 10, 20)

	fmt.Println(cap(s), len(s))

	s = append(s, 1, 2, 4, 5, 6, 7)
	fmt.Println(cap(s), len(s))

}
