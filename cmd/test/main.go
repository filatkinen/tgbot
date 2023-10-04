package main

import "fmt"

func main() {

	//ar := []int{1, 3, 4, 5, 6}

	ar1 := make([]int, 10)
	for i := 0; i < len(ar1); i++ {
		ar1[i] = i
	}
	fmt.Println(ar1[0:1])
	fmt.Println(ar1[0:2])

	fmt.Println(ar1[1:1])
	fmt.Println(ar1[1:2])

}
