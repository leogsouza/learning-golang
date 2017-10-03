package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		// Let the CPU execute other goroutines and come back at this point
		runtime.Gosched()
		fmt.Println(s)
	}
}

func main() {
	go say("world") // create a new goroutine
	say("hello")    // current goroutine
}
