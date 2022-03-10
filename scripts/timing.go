package main

import (
	"fmt"
	"time"
	"local/jgbaldwinbrown/siter/siter"
)

func main() {
	a := []int{5,6,7}
	iter := SliceRange(&a)
	for iter.Next() {
		fmt.Println(iter.Value())
	}

	iter = SliceRange(&a)
	for iter.Next() {
		ptr := iter.Ptr().(*int)
		*ptr++
	}
	fmt.Println(a)

	a2 := make([]int, 10000000)
	a3 := make([]int, 10000000)
	iter = SliceRange(&a2)
	i := 0
	start := time.Now()
	for iter.Next() {
		a3[i] = a2[i]
		i++
	}
	end := time.Now()
	fmt.Println(end.Sub(start))
}
