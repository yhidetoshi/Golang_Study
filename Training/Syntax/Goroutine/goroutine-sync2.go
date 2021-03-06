package main

import (
	"fmt"
	"sync"
	"time"
)

const N = 1000000000

func main() {
	// sync WaitGroupを生成

	wg := new(sync.WaitGroup)
	isFin := make(chan bool)
	wg.Add(2)
	start := time.Now()
	go hoge(isFin, wg)
	go fuga(isFin, wg)
	wg.Wait()
	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
	close(isFin)
}

var sumHoge int

func hoge(isFin chan bool, wg *sync.WaitGroup) {
	for i := 0; i < N; i++ {
		//fmt.Println("Goroutine hogeee +++")
		sumHoge++
	}
	wg.Done()
	isFin <- true
}

var sumFuga int

func fuga(isFin chan bool, wg *sync.WaitGroup) {
	for i := 0; i < N; i++ {
		//fmt.Println("Goroutine fugaaa ---")
		sumFuga++
	}
	wg.Done()
	isFin <- true
}
