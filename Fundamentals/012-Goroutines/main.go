package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 6; i++ {
		time.Sleep(time.Second)
		fmt.Println(from, ":", i)
	}
}

func main() {

	f("direct")

	go f("goroutine")

	go func(msg string) {
		// fmt.Println(msg)
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second)
			fmt.Println(msg, ":", i)
		}
	}("going")

	time.Sleep(time.Second * 5)
	fmt.Println("done")
}
