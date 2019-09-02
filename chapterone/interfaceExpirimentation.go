package chapterone

import (
	"fmt"
)

//I ...
type I interface {
	M()
}

//T ...
type T struct {
	S string
}

//M ...
func (t T) M() {
	fmt.Println(t.S)
}

//InterfaceOne ...
func InterfaceOne(implementation T) {
	var a I = implementation
	a.M()
}
