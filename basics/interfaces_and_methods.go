package basics

import "fmt"

type Doggo struct {
	name     string
	isHungry bool
}

// non-struct types can also have a methods (see iota_enum.go)

// just a function with receiver arg
// receiver should be available within package
func (d Doggo) Woof() {
	fmt.Println("Woof! I'm", d.name)
}

func WoofExplicit(d Doggo) {
	fmt.Println("Woof! I'm", d.name)
}

// this method doesn't change internal state
// only copy has passed
func (d Doggo) EatBadFood() {
	d.isHungry = false
}

// pass by pointer, now it can
func (d *Doggo) EatChappie() {
	d.isHungry = false
}

func interfacesAndMethods() {
	var bill Doggo = Doggo{}

	bill.name = "Bill"
	bill.Woof()

	WoofExplicit(bill)

	bill.isHungry = true
	bill.EatBadFood()
	fmt.Println("bill.isHungry (bad food) ==", bill.isHungry)
	bill.EatChappie()
	fmt.Println("bill.isHungry (chappie) ==", bill.isHungry)
}
