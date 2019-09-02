package chapterone

import "fmt"

//Vertex ...
type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

//PrintVertex :
func PrintVertex(x, y int) {
	fmt.Println(Vertex{y, x})
}

//ManipulateVertex ...
func ManipulateVertex(x, y int) {
	v := Vertex{x, y}
	v.X = 10
	fmt.Println(v)

	fmt.Println(v1, p, v2, v3)
}
