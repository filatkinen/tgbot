package main

import "fmt"

func main() {

	newopts()
	newopts(10, 2, 3)

}

func newopts(n ...int) {
	for _, v := range n {
		fmt.Println(v)
	}
}
