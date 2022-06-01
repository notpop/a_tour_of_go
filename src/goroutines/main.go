package main

import (
	"fmt"
	"time"
)

func waitOneSecond() {
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("1 second passed")
}

func waitTwoSecond() {
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("2 second passed")
}

func waitThreeSecond() {
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("3 second passed")
}

func main() {
	fmt.Println(time.Now())
	waitThreeSecond()
	go waitTwoSecond()
	go waitOneSecond()
	fmt.Println(time.Now())
}
