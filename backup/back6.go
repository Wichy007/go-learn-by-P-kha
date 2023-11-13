package main

import (
	"fmt"
	"time"
)

func Ping1S(c chan string) {
	time.Sleep(1 * time.Second)
	// fmt.Println("success form 1 s")

	c <- "success from 1 s"

}

func Ping5S(ch chan string) {
	time.Sleep(5 * time.Second)
	// fmt.Println("success form 5 s")

	ch <- "success from 5 s"

}

func back6() {
	start := time.Now()
	// time.Sleep(5 * time.Second)

	c := make(chan string)
	ch := make(chan string)

	go Ping1S(c)
	go Ping5S(ch)

	fromPing1S, fromPing5S := <-c, <-ch
	fmt.Println(fromPing1S)
	fmt.Println(fromPing5S)

	fmt.Println(time.Since(start))

}
