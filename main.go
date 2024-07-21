package main

import (
	"fmt"
	"math/rand"
	"time"
	"github.com/Georgery/go_sort/sort"
)

func main() {
	x := rand_int(10000)
	// fmt.Println(x)

	start := time.Now()
	sort.Bubble(x)
	this_took := time.Since(start)
	fmt.Println("Bubble Sort took:", this_took)
	// fmt.Println(x)
}

func rand_int(size int) []int {
	x := make([]int, 0, size)
	for i := 0; i < size; i++ {
		x = append(x, rand.Intn(100000))
	}
	return x
}


