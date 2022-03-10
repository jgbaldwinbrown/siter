package siter

import (
	"local/jgbaldwinbrown/xunsafe"
	"unsafe"
	"reflect"
)

type IndexedVals interface {
	Value(int) interface{}
	Len() int
}

type IndexedPtrs interface {
	Ptr(int) interface{}
	Len() int
}

type Iter interface {
	Next() bool
	Value() interface{}
}

type MutIter interface {
	Next() bool
	Ptr() interface{}
}

type SliceIndexed struct {
	SlicePtr interface{}
	Slicer *xunsafe.Slice
	Start unsafe.Pointer
	Length int
}

type IndexedIter struct {
	Indexed IndexedVals
	Pos int
}

type IndexedMutIter struct {
	Indexed IndexedPtrs
	Pos int
}

func SliceIndex(sliceptr interface{}) (idx SliceIndexed) {
	idx.SlicePtr = sliceptr
	idx.Slicer = xunsafe.NewSlice(reflect.ValueOf(sliceptr).Elem().Type())
	idx.Start = unsafe.Pointer(reflect.ValueOf(idx.SlicePtr).Pointer())
	idx.Length = idx.Slicer.Len(idx.Start)
	return idx
}

func (i *SliceIndexed) Value(pos int) interface{} {
	return i.Slicer.ValueAt(i.Start, pos)
}

func (i *SliceIndexed) Ptr(pos int) interface{} {
	return i.Slicer.AddrAt(i.Start, pos)
}

func (i *SliceIndexed) Len() int {
	return i.Length
}

func IndexedRange(i IndexedVals) (iter IndexedIter) {
	iter.Indexed = i
	iter.Pos = -1
	return
}

func (i *IndexedIter) Next() bool {
	i.Pos++
	return i.Pos < i.Indexed.Len()
}

func (i *IndexedIter) Value() interface{} {
	return i.Indexed.Value(i.Pos)
}

func SliceRange(sliceptr interface{}) IndexedIter {
	indexed := SliceIndex(sliceptr)
	return IndexedRange(&indexed)
}

func IndexedMutRange(i IndexedPtrs) (iter IndexedMutIter) {
	iter.Indexed = i
	iter.Pos = -1
	return
}

func (i *IndexedMutIter) Next() bool {
	i.Pos++
	return i.Pos < i.Indexed.Len()
}

func (i *IndexedMutIter) Ptr() interface{} {
	return i.Indexed.Ptr(i.Pos)
}

func SliceMutRange(sliceptr interface{}) IndexedMutIter {
	indexed := SliceIndex(sliceptr)
	return IndexedMutRange(&indexed)
}
