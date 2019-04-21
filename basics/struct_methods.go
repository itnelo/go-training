package basics

import "fmt"

type Doggo struct {
	name     string
	isHungry bool
}

// non-struct types can also have a methods (see iota_enum.go)

// convention: either all pointer receivers or all value receivers, but not both

// just a function with receiver arg
// receiver should be available within package
func (d Doggo) Woof() {
	fmt.Println("Woof! I'm", d.name, ", am I hungry?", d.isHungry)
}

func WoofExplicit(d Doggo) {
	fmt.Println("Woof! I'm", d.name, ", am I hungry?", d.isHungry)
}

// this method doesn't change internal state
// only copy has passed
func (d Doggo) EatBadFood() {
	// modifying a local copy
	// will not be modified outside this visibility context
	d.isHungry = false
}

// pass by pointer, now it can
func (d *Doggo) EatChappie() {
	d.isHungry = false
}

func structMethods() {
	var bill Doggo = Doggo{}

	bill.name = "Bill"
	bill.Woof()

	WoofExplicit(bill)

	bill.isHungry = true
	bill.EatBadFood()
	fmt.Println("bill.isHungry (bad food) ==", bill.isHungry)
	bill.EatChappie()
	fmt.Println("bill.isHungry (chappie) ==", bill.isHungry)

	// difference between explicit functions - methods with receiver
	// accept both copy and pointer
	var samPtr *Doggo = &Doggo{"sam", true}
	samPtr.Woof()
	samPtr.EatBadFood()
	samPtr.Woof()
	samPtr.EatChappie()
	samPtr.Woof()

	// also work in reverse direction (indirection)
	// samPtr becomes (*samPtr) if needed
}
