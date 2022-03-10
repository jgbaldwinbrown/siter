# Siter

A generic slice iterator library that _does not_ use Go 1.18's generics.

## Introduction

I am as happy as anyone that Go is getting generics now, finally, in version
1.18. It will make generic libraries both possible and sensible in a way that
they weren't before. In honor of the new generics, though, I wanted to see how
far Go's old methods for generics (reflection and unsafe pointers) could be
taken. Using the excellent library _xunsafe_, which uses unsafe pointers to do
lightning-fast reflection operations, I was able to write a generic library
that generates an iterator over slice elements for slices of any kind.

Interestingly, I've seen a lot of people request a feature like this, only for
someone to say it's impossible in Go without proper generics support. This
probably says more about the Go community's general distaste for reflection and
unsafe behavior than it does about Go itself. I'm inclined to agree that
memory- and type-unsafe behavior should be avoided, but this library shows how
it can be safely encapsulated such that library users can avoid any use of
_unsafe_, while still gaining the benefits of its use.

All of this to say, this library is not intended for real use, but is more
a proof of concept. Generics are not available everywhere yet, but they will
be soon, and once they are, libraries like this will have no purpose. Still,
it makes me wonder why people waited so long to introduce libraries like this,
given that they can be fast and (relatively) safe. Maybe it says something about
the baked-in dogma in most programming language cultures. Regardless, I'm looking
forward to a generic future with Go!

## Installation

go get github.com/jgbaldwinbrown/siter/siter

## Usage

Here's a simple example of siter in action:

```go
package main

import (
	"fmt"
	"siter"
)

func main() {
	// A typical slice of ints. Note that this slice could contain any kind
	// of value -- siter doesn't care.
	a := []int{5,6,7}

	// Generate an iterator that will iterate over all of the elements
	// of a. Iterators have three methods, Next(), Value(), and Ptr().
	// These are mostly self-explanatory, with Ptr() giving a pointer
	// to the current element for setting purposes.
	iter := siter.SliceRange(&a)

	// Loop over the iterator and print all of the values in the slice.
	for iter.Next() {
		fmt.Println(iter.Value())
	}

	// Loop over the iterator and modify the values inside it.
	iter = siter.SliceRange(&a)
	for iter.Next() {
		ptr := iter.Ptr().(*int)
		*ptr++
	}
	fmt.Println(a)

}
```

## Benchmarks

Here are some quick benchmarks from `scripts/timing.go`:

| Method | Time | Normalized time |
| --- | --- | --- |
| Non-generic | 21.509735ms | 1.0 |
| siter | 173.209838ms | 8.055 |
| reflection | 1.911873678s | 88.88 |

The short version is that siter is about 8 times slower than totally
non-generic code, but reflection is 88 times slower. As far as generic
code goes, you can't beat siter on performance.
