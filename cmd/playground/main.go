package main

import (
	"fmt"
	"github.com/YaroslavRostanets/memoizer"
	"time"
)

func main() {
	mem := cache.New()

	mem.Set("qwerty", 42, time.Second)
	mem.Set("qwerty1", 85, time.Second)
	fmt.Println(mem.Get("qwerty"))
	fmt.Println(mem.Get("qwerty1"))
	time.Sleep(10 * time.Second)
	fmt.Println(mem.Get("qwerty"))
}
