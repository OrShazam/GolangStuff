package main

// demonstration of concurrency

import (
	"fmt"
	"time"
)
func f (c chan int){
	for {
		time.Sleep(2 * time.Second);
		count := <-c + 1
		c <- count;
		fmt.Println("f ->", count);
	}
}
func main() {
	c := make(chan int);
	go f(c);
	c <- 0;
	for {
		count := <-c + 1
		c <- count;
		fmt.Println("main ->", count);
		time.Sleep(2 * time.Second);
	}	
}