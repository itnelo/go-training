package basics

import (
	"fmt"
	"unsafe"
)

// no keywords, just a method lookup at compile time
// interface is a (value, type) container
// also can be considered as a holder of pointer to a method in virtual table
// and a pointer to concrete data instance (receiver)

// Go uses method tables (like C++, Java), but computes them at runtime

// conventions: reference *receiver, nil handling

// https://research.swtch.com/interfaces
// http://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go

type Talker interface {
	talk() string
}

type Parrot struct {
	name      string
	needWater bool
}

func (talker Parrot) talk() string {
	defer func() {
		// we doesn't change a real talker instance here, just a local copy
		talker.needWater = true
	}()

	return "parrot talks"
}

type Badger struct {
	name      string
	needWater bool
}

func (talker *Badger) talk() string {
	if nil == talker {
		return "<nil>"
	}

	defer func() {
		talker.needWater = true
	}()

	return "badger talks"
}

func nilUnderlyingValue() {
	var talker Talker

	// panic: runtime error: invalid memory address or nil pointer dereference
	// talker.talk()

	var badger *Badger

	talker = badger

	fmt.Printf("(%v, %T)\n", talker, talker)

	// "<nil>"
	talker.talk()

	// another way
	fmt.Println("An explicit method call by pointer with a receiver arg:", (*Badger).talk(nil))
}

// all types satisfy the empty interface
func takeAnyValue(i interface{}) {
	fmt.Printf("An empty interface: (%#v, %T)\n", i, i)
}

func emptyInterface() {
	// can hold any tuple (value, type) at runtime
	var i interface{}

	i = 5
	takeAnyValue(i)

	i = "now i am a string holder"
	takeAnyValue(i)

	i = Badger{"moves like a badger", false}
	takeAnyValue(i)
}

func interfaces() {
	// (value, type)
	var talker Talker

	parrot := Parrot{"ben", false}

	talker = parrot
	talker.talk()

	talkersNeedWater := ((*Parrot)(unsafe.Pointer(&talker))).needWater
	fmt.Printf("memaddr parrot(%t)(%p) talker(%t)(%p)\n", parrot.needWater, &parrot, talkersNeedWater, &talker)

	badger := Badger{"kim", false}

	// Badger does not implement Talker (talk method has pointer receiver)
	// talker = badger

	// ok
	talker = &badger
	talker.talk()

	// looks like a pointer to separate instance
	// we can see some internals of underlying struct with "Badger" glasses
	var трешПоинтер = (*Badger)(unsafe.Pointer(&talker))

	// %p &badger == %p *&talker
	// &talker - memaddr of "interface container", that holds (value, type)
	fmt.Printf("memaddr badger(%t)(%p) talker(%t)(%p)\n", badger.needWater, &badger, трешПоинтер.needWater, &talker)

	трешПоинтер.needWater = true
	fmt.Printf("memaddr badger(%t)(%p) talker(%t)(%p)\n", badger.needWater, &badger, трешПоинтер.needWater, &talker)

	fmt.Printf("sizeof badger(%d) talker(%d)\n", unsafe.Sizeof(badger), unsafe.Sizeof(talker))

	// memaddr parrot(false)(0xc000130040) talker(false)(0xc0001140a0)
	// memaddr badger(true)(0xc000130080) talker(false)(0xc0001140a0)
	// memaddr badger(true)(0xc000130080) talker(true)(0xc0001140a0)
	// sizeof badger(24) talker(16)

	nilUnderlyingValue()

	emptyInterface()
}
