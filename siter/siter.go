package main

import (
	"fmt"
	"local/jgbaldwinbrown/xunsafe"
	"unsafe"
	"reflect"
	"time"
)

type Iter interface {
	Next() bool
	Value() interface{}
}

type SliceIter struct {
	SlicePtr interface{}
	Slicer *xunsafe.Slice
	Start unsafe.Pointer
	Len int
	Pos int
}

func SliceRange(sliceptr interface{}) (iter SliceIter) {
	iter.SlicePtr = sliceptr
	iter.Slicer = xunsafe.NewSlice(reflect.ValueOf(sliceptr).Elem().Type())
	iter.Start = unsafe.Pointer(reflect.ValueOf(iter.SlicePtr).Pointer())
	iter.Len = iter.Slicer.Len(iter.Start)
	iter.Pos = -1
	return iter
}

func (i *SliceIter) Value() interface{} {
	return i.Slicer.ValueAt(i.Start, i.Pos)
}

func (i *SliceIter) Ptr() interface{} {
	return i.Slicer.AddrAt(i.Start, i.Pos)
}

func (i *SliceIter) Next() bool {
	i.Pos++
	return i.Pos < i.Len
}

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
