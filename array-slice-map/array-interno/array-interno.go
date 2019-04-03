package main

import "fmt"

func main() {
	s1 := make([]int, 10, 20)
	s2 := append(s1, 1, 2, 0)
	s1[6] = 12
	fmt.Println(s1, s2)
}
