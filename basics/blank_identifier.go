package basics

import (
	"fmt"

	// side effect import / blank import
	// https://golang.org/doc/effective_go.html#blank_import
	_ "net/http/pprof"
)

// conventions:
// - blank import only in main/tests

type Array5 [5]int

func blankIdentifier() {
	// https://golang.org/doc/effective_go.html#blank
	//s := [...]int{1, 2, 3, 4, 5}
	s := Array5{1, 2, 3, 4, 5}

	// "s" is under development / TODO mark
	_ = s

	// omitting an index num
	for _, value := range s {
		fmt.Println(value)
	}
}
