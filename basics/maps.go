package basics

import (
	"fmt"
)

type Cat struct {
	age int
}

func (c Cat) Meow() {
	fmt.Println("A cat meows: ", c.age)
}

var (
	cats map[string]Cat
)

func maps() {
	cats = make(map[string]Cat)
	fmt.Printf("Initialized map %v with len %v\n", cats, len(cats))

	// oh, this works.
	cats["Aylin"].Meow() // A cat meows: 0

	cats["Aylin"] = Cat{age: 2}
	cats["Aylin"].Meow()

	fmt.Printf("%v with len %v\n", cats, len(cats))

	// like a struct, keys is required
	var catsByLiteral map[string]Cat = map[string]Cat{
		"Aylin": Cat{age: 3},
		"Jane":  Cat{age: 10}, // last comma is required
	}

	fmt.Printf("%v with len %v\n", catsByLiteral, len(catsByLiteral))

	// omitting top-level type
	catsByFluentLiteral := map[string]Cat{
		"Aylin": {4},
		"Jane":  {11},
	}

	fmt.Printf("%v with len %v\n", catsByFluentLiteral, len(catsByFluentLiteral))
	catsByFluentLiteral["Jane"].Meow()

	// mutation, retrieve
	cats["Bill"] = Cat{age: 16}
	var billTheCat Cat = cats["Bill"]
	fmt.Printf("Added Bill: %v with len %v\n", cats, len(cats))
	billTheCat.Meow()

	// deletion
	delete(cats, "Bill")
	cat, isBillExists := cats["Bill"]
	fmt.Println("Is Bill exists:", isBillExists, cat)
}