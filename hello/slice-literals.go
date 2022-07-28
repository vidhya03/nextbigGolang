package main

import "fmt"

func main() {
	fmt.Println("hello universe")
	array := []int{1, 2, 3, 4, 5}
	fmt.Println(array)
	boo := []bool{true, true, true, true, true}
	fmt.Println(boo)
	mystruct := []struct {
		i int
		b bool
	}{

		{i: 1, b: true},
		{2, true},
		{b: true, i: 3},
	}
	fmt.Println(mystruct)

}
