package basics

import (
	"fmt"
)

// maps are references to runtime type

// we cannot take pointer directly to map element
// because
// 1) map moves element to another memory addr while grown,
// so it can lead to dangling pointer (restricted by Go design)
// map is not a slice (doesn't have an underlying array)

// 1.11: map can't be read-only, need to be careful while passing.

type Cat struct {
	age int
}

func (c Cat) Meow() {
	fmt.Println("A cat meows: ", c.age)
}

var (
	cats map[string]Cat
)

// data will be changed because values are stored in heap.
func changeMap(m map[string]Cat) {
	m["Bob"] = Cat{age: 8}
}

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

	// changing within a function
	changeMap(cats)
	fmt.Println(cats)
}
