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

	// todo: slice copy
	// todo: delete slice elem
	// todo: slice comparison
	// https://medium.com/rungo/the-anatomy-of-slices-in-go-6450e3bb2b94
	// https://github.com/golang/go/wiki/SliceTricks
}
