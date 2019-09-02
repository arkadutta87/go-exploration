package chapterone

import (
	"fmt"
)

//Add x , y => x+y
func Add(x, y int) int {
	return x + y
}

//Swap : x, y => y, x
func Swap(x, y string) (string, string) {
	return y, x
}

//Split : sum => x, y
func Split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x

	return
}

//SumXNaturalNumbers : return 0 +1 +... +x
func SumXNaturalNumbers(x int) int {
	sum := 0
	for i := 0; i <= x; i++ {
		sum += i
	}

	return sum
}

//PointerTesting : testing with pointers
func PointerTesting() {
	i, j := 42, 2701

	var p *int
	p = &i

	fmt.Println(p, j)
	fmt.Println(*p)
	*p = 21
	fmt.Println(i, *p)

	p = &j
	*p = *p / 37
	fmt.Println(j)

}
