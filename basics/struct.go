package basics

import (
	"fmt"
)

type Point2D struct {
	X int
	Y int
}

// Go: pass by value forever

func localStructCopy(p Point2D) {
	p.X = 7
}

func pointerToStruct(p *Point2D) {
	p.X = 8
}

func pointerIsALocalCopyAndCannotBeRedirected(p *Point2D) {
	p = &Point2D{3, 3}
}

func Structures() {
	// same as Point2D{X: 5, Y: 25}, Point2D{Y: 25, X: 5}
	// Point2D{X: 5, 25} is invalid
	// Point2D{} with default values for fields
	point := Point2D{5, 25}
	pointPtr := &point

	fmt.Printf("Point2D: %v, Ptr: %p\n", point, pointPtr)

	localStructCopy(point)
	fmt.Println("Local copy has created so any modifications are not saved: ", point)

	pointerToStruct(pointPtr)
	fmt.Println("Modification by address in heap, so changes are applied: ", point)

	pointerIsALocalCopyAndCannotBeRedirected(pointPtr)
	fmt.Println("Pointer itself is a local copy and cannot be redirected: ", *pointPtr)

	(*pointPtr).X = 100
	fmt.Println("(*pointPtr).X = 100, result:", pointPtr)

	// is just a syntactic sugar, an implicit dereferencing
	pointPtr.X = 200
	pointPtr.Y = 1e9
	fmt.Println("implicit dereferencing pointPtr.x = 200, result:", pointPtr)

	// new statement
	//var ptr2 = new(Point2D)
}
