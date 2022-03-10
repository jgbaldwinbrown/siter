package main

import (
	"fmt"
	"time"
	"local/jgbaldwinbrown/siter/siter"
	"reflect"
)

type SliceIter struct {
	Slice interface{}
	SliceVal reflect.Value
	SliceValVal reflect.Value
	SliceValValVal reflect.Value
	Started bool
}

func (si *SliceIter) Value() interface{} {
	return si.SliceVal.Elem().Elem().Index(0).Interface()
}

func (si *SliceIter) Next() bool {
	if si.SliceValValVal.Len() < 1 {
		return false
	}
	if si.Started == false {
		si.Started = true
	} else {
		si.SliceValVal.Set(si.SliceValValVal.Slice(1, si.SliceValValVal.Len()))
	}
	si.SliceValValVal = si.SliceValVal.Elem()
	if si.SliceValValVal.Len() < 1 {
		return false
	}
	return true
}

func MakeSliceIter(slice interface{}) SliceIter {
	out := SliceIter{}
	out.Slice = slice
	out.SliceVal = reflect.ValueOf(&out.Slice)
	out.SliceValVal = out.SliceVal.Elem()
	out.SliceValValVal = out.SliceVal.Elem().Elem()
	out.Started = false
	return out
}

type MapIter struct {
	MI *reflect.MapIter
}

type KeyValuePair struct {
	Key interface{}
	Value interface{}
}

func MakeMapIter(amap interface{}) MapIter {
	av := reflect.ValueOf(amap)
	return MapIter{av.MapRange()}
}

func (m *MapIter) Next() bool {
	return m.MI.Next()
}

func (m *MapIter) Value() interface{} {
	return KeyValuePair{m.MI.Key().Interface(), m.MI.Value().Interface()}
}

func time_siter() {
	a := []int{5,6,7}
	idx := siter.SliceIndex(&a)
	iter := siter.IndexedRange(&idx)
	for iter.Next() {
		fmt.Println(iter.Value())
	}

	iter2 := siter.IndexedMutRange(&idx)
	for iter2.Next() {
		ptr := iter2.Ptr().(*int)
		*ptr++
	}
	fmt.Println(a)

	a2 := make([]int, 10000000)
	a3 := make([]int, 10000000)
	idx2 := siter.SliceIndex(&a2)
	iter3 := siter.IndexedRange(&idx2)
	i := 0
	start := time.Now()
	for iter3.Next() {
		a3[i] = iter3.Value().(int)
		i++
	}
	end := time.Now()
	fmt.Println(end.Sub(start))


}


func time_fast_and_slow() {
	a := []int{5, 6, 7}
	b := []string{"five", "six", "seven"}
	ai := MakeSliceIter(a)
	for ai.Next() {
		fmt.Println(ai.Value())
	}
	bi := MakeSliceIter(b)
	for bi.Next() {
		fmt.Println(bi.Value())
	}

	c := map[string]float64 {"banana": 9.8, "apple": 11.7}
	ci := MakeMapIter(c)
	for ci.Next() {
		fmt.Println(ci.Value())
	}

	long := make([]int, 10000000)
	newlong := make([]int, 10000000)
	start1 := time.Now()
	for i, _ := range long {
		newlong[i] = long[i] + 1
	}
	end1 := time.Now()

	newlong2 := make([]int, 10000000)
	start2 := time.Now()
	li := MakeSliceIter(long)
	i := 0
	for li.Next() {
		newlong2[i] = li.Value().(int)
		i++
	}
	end2 := time.Now()
	fmt.Println(end1.Sub(start1))
	fmt.Println(end2.Sub(start2))
}

func main() {
	time_siter()
	time_fast_and_slow()
}
