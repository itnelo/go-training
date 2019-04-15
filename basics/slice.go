package basics

import (
	"fmt"
	"math/rand"
	"reflect"
)

// variadic
func sum(numbers ...int) (sum int) {
	defer func() {
		if sum > 10 {
			fmt.Println("sum is called (latest call)")
		}
	}()

	for _, number := range numbers {
		defer func(n int) {
			fmt.Println("sum:", sum, "+", n)

			sum += n
		}(number)
	}

	return
}

func shortArrayLiteral() {
	s := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("%v (%v), len(s) == %v, cap(s) == %v\n", s, reflect.TypeOf(s), len(s), cap(s))
}

func sliceAppend() {
	a := make([]int, 5, 5)

	rand.Seed(293)
	for i := 0; i < 5; i++ {
		a[i] = rand.Intn(11)
	}

	fmt.Printf("%v, len(a) == %v, cap(a) == %v, addr == %p\n", a, len(a), cap(a), &a)

	a = append(a, 1)
	fmt.Printf("%v, len(a) == %v, cap(a) == %v, addr == %p\n", a, len(a), cap(a), &a)
	a = append(a, 3, 300, 3000) // multiple append
	fmt.Printf("%v, len(a) == %v, cap(a) == %v, addr == %p\n", a, len(a), cap(a), &a)

	// [2,5)
	var as []int = a[2:5]
	fmt.Println("a[2:5] slice:", as)

	// acts like a refernce to an array
	as[0] = 100 // a[2] = 100
	fmt.Println("as[0] = 100, a now:", a)
}

func anonymousStructSlice() {
	structSlice := []struct {
		x, y int
	}{
		{1, 5},
		{100, 200},
		{400, 450}, // last comma is not an error
	}

	fmt.Println("Slice of anonymous struct instances:", structSlice)
}

func omitBounds(slice []string) {
	sc1 := slice[:]
	fmt.Println("Omitting both low and high bounds:", sc1)

	// slice[0:]
	// slice[:5]
}

func slices() {
	// [5]int: array
	// []int: slice
	// slicing doesn't change an actual array under it
	// only assignment by index does

	shortArrayLiteral()

	sliceAppend()

	// as an argument to variadic function
	b := []int{1, 2, 3, 4, 5}
	fmt.Println("Unpacking a slice:", b, "sum:", sum(b...))
	fmt.Println("Call to sum() without a slice:", sum(1, 2, 3, 4, 5))

	anonymousStructSlice()

	omitBounds([]string{"one", "two", "three"})

	// not a nil slice
	var bb = b[:0]
	bb = nil
	fmt.Println("empty but not a nil slice: ", bb, "len:", len(bb), "cap:", cap(bb), "nil:", bb == nil)
	// a nil slice
	var c []int
	fmt.Println("a nil slice: ", c, "len:", len(c), "cap:", cap(c), "nil:", c == nil)

	// dynamically sized arrays
	var d []int = make([]int, 2, 7)
	fmt.Println("Dynamically sized array:", d)

	// https://medium.com/rungo/the-anatomy-of-slices-in-go-6450e3bb2b94
	// https://github.com/golang/go/wiki/SliceTricks
	// slice copy
	dst := make([]int, 3)
	var howMuchCopied int = copy(dst, b[1:4])
	dst[0] = 1000
	fmt.Printf(
		"Copied %v elements from b(%v, len: %v, cap: %v)[%p] to dst(%v, len: %v, cap: %v)[%p]\n",
		howMuchCopied,
		b, len(b), cap(b), &b,
		dst, len(dst), cap(dst), &dst)

	// delete slice elem
	dst = append(dst[:1], dst[2:]...)
	fmt.Println("Deleted element at index 1:", dst)

	dst2 := []int{1, 2, 3, 4, 5, 6, 7}
	// dst2[2:] == 4, 5, 6, 7
	// dst2[3:] == 5, 6, 7
	// copy -> 5, 6, 7, 7
	// copy[:3+3] -> 1, 2, 3, 5, 6, 7
	dst2 = dst2[:3+copy(dst2[3:], dst2[4:])]
	//dst2[len(dst2)-1] = nil // to eliminate leak (pointer/struct with underlying pointer)
	fmt.Println("Deleted element at index 3:", dst2)

	// slice comparison
	// with nil only
	// compare by range and O(n) loop for each element
}
