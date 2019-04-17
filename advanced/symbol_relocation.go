package advanced

import (
	_ "unsafe"
)

//go:linkname RelocatedForLoop github.com/itnelo/go-training/basics.forloop
func RelocatedForLoop()

func symbolRelocation() {
	RelocatedForLoop()
}
