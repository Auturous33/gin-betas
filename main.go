package main

import (
	"fmt"
	"time"
)

// it begins
func main() {
	startTime := time.Now()
	time.Sleep(10 * time.Second)
	fmt.Println("hello, gin-betas!")
	fmt.Printf("cost %f seconds\n", time.Since(startTime).Seconds())
}
