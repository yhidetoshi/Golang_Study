package main

import "fmt"

func main() {
	ch := make(chan string, 3)
	ch <- "Apple"
	fmt.Println(len(ch))
	fmt.Println(cap(ch))
}
