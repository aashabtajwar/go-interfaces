// EMPTY INTERFACES
/*
* An interface with zero methods is called an empty interface
* Since empty interface has zero methods, all types implement
* this interface implicitly
* The Println function from the fmt package accepts different
* types of values as arguments because of empty interface
 */

package main

import "fmt"

type MyString string

type Rect struct {
	width  float64
	height float64
}

// explain function which accepts an empty interface and prints
// the dynamic value and type of the interface
func explain(i interface{}) {
	fmt.Printf("Value given to this function is of type '%T' with value %v", i, i)
}

func main() {
	ms := MyString("Cyberpunk game")
	r := Rect{5.5, 4.5}
	explain(ms)
	explain(r)
}
