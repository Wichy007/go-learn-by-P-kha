package main

import "fmt"

func back3() {
	fmt.Println("hello world")

	// arr := []int{2,4,4,5}
	arr := make([]int, 4)
	arr[0] = 30
	fmt.Println(arr)

	txt := "today is sundat"
	fmt.Println(txt[0:5])
	fmt.Println(txt[0:1])
	fmt.Println(len(txt))
}
