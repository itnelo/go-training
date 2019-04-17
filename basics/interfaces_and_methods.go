package basics

import "fmt"

type Doggo struct {
	name string
}

// just a function with receiver arg
func (d Doggo) Woof() {
	fmt.Println("Woof! I'm", d.name)
}

func WoofExplicit(d Doggo) {
	fmt.Println("Woof! I'm", d.name)
}

func interfacesAndMethods() {
	var bill Doggo = Doggo{}

	bill.name = "Bill"
	bill.Woof()

	WoofExplicit(bill)
}
