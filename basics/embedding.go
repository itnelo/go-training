package basics

import (
	"fmt"
)

// visibility is on a per package only

// https://golang.org/doc/effective_go.html#embedding

type Кот interface {
	talk()
}

type Пёс interface {
	// duplicate method talk
	//talk()
	talkAsDog()
}

// only interfaces can be embedded within interfaces
type Котопёс interface {
	Кот
	Пёс
}

type Мурзик struct {
	name string
}

func (кот *Мурзик) talk() {
	fmt.Printf("Кот(%v)\n", кот)
}

type Бобик struct {
	name string
}

func (пёс *Бобик) talkAsDog() {
	fmt.Println("Пёс")
}

type Мурзебобик struct {
	// direct embedding
	// we don't need to explicitly write boilerplate code for methods forwarding
	// we can target a specific struct within methods by similar field names
	Мурзик       // *Мурзик requires initialization
	dog    Бобик // field name requires explicit forwarding method, *Бобик requires initialization
	// forwarding methods is provided implicitly
	// the receiver will be a concrete inner struct: Мурзик or Бобик, not Мурзебобик itself
}

func (котопёс *Мурзебобик) talkAsDog() {
	if nil == котопёс {
		return
	}

	fmt.Println("forwarding talkAsDog() from outer to inner struct")

	котопёс.dog.talkAsDog()
}

func (котопёс *Мурзебобик) setDogName(name string) {
	котопёс.dog.name = name
}

func (котопёс *Мурзебобик) getDogName() string {
	return котопёс.dog.name
}

// talker *Котопёс => talker.talk undefined (type *Котопёс is pointer to interface, not interface)
func ask(talker interface{}) {
	fmt.Printf("(%#v, %T)\n", talker, talker)

	if котопёс, ok := talker.(Котопёс); ok {
		котопёс.talk()
		котопёс.talkAsDog()
	}
}

func embedding() {
	var котопёс = &Мурзебобик{}

	ask(котопёс)

	котопёс.setDogName("new dog name2")
	fmt.Println(котопёс.getDogName())

	// (&basics.Мурзебобик{Мурзик:basics.Мурзик{name:""}, dog:basics.Бобик{name:""}}, *basics.Мурзебобик)
	// Кот(&{})
	// forwarding talkAsDog() from outer to inner struct
	// Пёс
	// new dog name2
}
