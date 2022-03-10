package siter

import (
	"local/jgbaldwinbrown/xunsafe"
	"unsafe"
	"reflect"
)

type Iter interface {
	Next() bool
	Value() interface{}
}

type MutIter interface {
	Next() bool
	Ptr() interface{}
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
