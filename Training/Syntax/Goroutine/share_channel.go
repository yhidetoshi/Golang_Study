package main

import (
	"fmt"
)

func reciver(ch <-chan int) {
	for {
		i := <-ch
		fmt.Println(i)
	}
}

func main() {
	ch := make(chan int)
	go reciver(ch)

	i := 0
	for i < 10 {
		ch <- i
		i++
	}

}
